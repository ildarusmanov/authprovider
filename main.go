package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Printf("[*] Starting")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Printf("[*] Ready")

	<-stop

	log.Printf("[*] Finished")
}
