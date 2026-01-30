package handlers

import (
	"log"
	"net/http"

	"github.com/musishere/microservices/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}

	if r.Method == http.MethodPut {
		// expect the id in the URI
		p := r.URL.Path
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handler Post Products")

	prod := &data.Product{}

	// Attempt to decode the product data from the request body into the prod object
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}
