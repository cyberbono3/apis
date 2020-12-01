package main

import (
	"net"
	"os"
	"log"

	protos "github.com/cyberbono3/apis/currency/protos/currency"
	"github.com/cyberbono3/apis/currency/server"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
)

func main() {

	

	log := log.New(os.Stdout, "currency services", log.LstdFlags)

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs,cs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Println("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
