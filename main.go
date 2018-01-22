package main

import (
	"flag"
	"github.com/ildarusmanov/authprovider/grpcserver"
	"github.com/ildarusmanov/authprovider/providers"
	"github.com/ildarusmanov/authprovider/services"
	"log"
	"os"
	"os/signal"
)

var rvToken = flag.String("token", "request-validator-token", "Secret token")

func main() {
	flag.Parse()

	log.Printf("[*] Starting")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Printf("[*] Ready")

	srv, err := grpcserver.StartServer(
		services.CreateNewRequestValidator(*rvToken),
		providers.CreateNewMemoryTokenProvider(),
	)

	if err != nil {
		log.Fatalf("Can not start grpc server, error: %s", err)
	}

	defer srv.GracefulStop()

	<-stop

	log.Printf("[*] Finished")
}
