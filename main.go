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

	v := services.CreateNewRequestValidator(*rvToken)
	p := providers.CreateNewMemoryTokenProvider()
	grpcserver.StartServer(v, p)

	<-stop

	log.Printf("[*] Finished")
}
