package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "chaos-io/microserver/currency/protos/currency"
	"chaos-io/microserver/currency/server"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	c := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, c)
	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
