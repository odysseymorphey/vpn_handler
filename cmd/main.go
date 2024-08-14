package main

import (
	"log"
	"vpn-handler/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal(err)

	}

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}
}
