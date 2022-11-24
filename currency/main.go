package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/chaos-io/microserver/currency/protos/currency"
	"github.com/chaos-io/microserver/currency/server"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	c := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, c)
	reflection.Register(gs)

	log.Info("Starting server on port 9093")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9093))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
