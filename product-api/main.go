package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	protos "github.com/chaos-io/microserver/currency/protos/currency"
	"github.com/chaos-io/microserver/product-api/data"
	"github.com/chaos-io/microserver/product-api/handlers"

	gohandlers "github.com/gorilla/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var bindAddress = flag.String("BIND_ADDRESS", ":9090", "Bind address for the server")

func main() {
	flag.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	conn, err := grpc.Dial("localhost:9093", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create client
	cc := protos.NewCurrencyClient(conn)

	// create the handlers
	pl := handlers.NewProducts(l, v, cc)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", pl.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", pl.GetProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", pl.UpdateProduct)
	putRouter.Use(pl.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", pl.AddProduct)
	postRouter.Use(pl.MiddlewareValidateProduct)

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/products/{id:[0-9]+}", pl.DeleteProduct)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	doc := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", doc)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create new server
	server := &http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      cors(sm),          // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time to connections using TCP Keep-Alive
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		l.Println("Starting server on port 9090")
		if err := server.ListenAndServe(); err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	// block until a signal is received
	sig := <-sigChan
	l.Printf("Received terminate, graceful shutdown %s\n", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
