package main

import (
	"log"
	"net/http"
	"os"

	"github.com/musishere/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc("/", nil)

	http.ListenAndServe(":9090", nil)
}
