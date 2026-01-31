package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("GoodBye world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Oops"))
		return
	}

	fmt.Fprintf(rw, "Hello from request writer %s", d)
}
