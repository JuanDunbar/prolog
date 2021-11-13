package main

import (
	"log"

	"github.com/juandunbar/prolog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":27088")
	log.Fatal(srv.ListenAndServe())
}
