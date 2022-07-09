package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Server(r *mux.Router, address string) http.Server {

	Server := http.Server{
		Addr:              address,
		Handler:           r,
		ReadTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 1 * time.Minute,
		MaxHeaderBytes:    1 << 20,
	}
	return Server
}

func main() {

	r := Router()
	server := Server(r, "0.0.0.0:8001")
	server.ListenAndServe()
	log.Println("running server....")

}
