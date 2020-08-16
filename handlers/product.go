package handlers

import (
	"encoding/json"
	"github.com/ashish19912009/anotherProject/data"
	"net/http"
	"log"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter,h *http.Request){
	lp := data.GetProducts()
	d,err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal the product list", http.StatusInternalServerError)
	}
	rw.Write(d)
}