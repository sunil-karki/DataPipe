package handlers

// curl localhost:9090/3 -XPUT | jq  -->  Replaces list that has Id 3
// curl localhost:9090/3 -XPUT -d '{"name": "sunilkarki", "description": "student", "price": 25, "sku": "skuee"}' | jq
// curl localhost:9090/ -X POST -d '{"name": "sunilkarki", "description": "student", "price": 25, "sku": "skuee"}'
// curl localhost:9090/ -X POST -d '{"fileid": 6, "position": 6, "filename": "sunilkarki", "description": "student", "filedate": "2020-02-04", "source": "ABC"}'

import (
	"log"
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
	// "github.com/nicholasjackson/building-microservices-youtube/product-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetData returns the data from the data source
func (p *Products) GetData(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Data")

	// // fetch the products from the datastore
	// lp := data.GetProducts()
	lp := data.GetData()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct add products to the list
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// prod := &data.Product{}

	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	// }

	// data.AddProduct(prod)
	// p.l.Printf("Prod: %#v", prod)

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// UpdateProducts updates the product when given an ID
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	// This line is for doing ...
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	// When Gorilla not used, following implementation were done..
	// p.l.Println("Handle PUT Product")

	// prod := &data.Product{}

	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	// }

	// err = data.UpdateProduct(id, prod)
	// if err == data.ErrProductNotFound {
	// 	http.Error(rw, "Product not found", http.StatusNotFound)
	// 	return
	// }

	// if err != nil {
	// 	http.Error(rw, "Product not found", http.StatusInternalServerError)
	// 	return
	// }
}

// KeyProduct added Context method ...
type KeyProduct struct{}

// MiddlewareValidateProduct gets executed before Handler.
// func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		prod := data.Product{}

// 		err := prod.FromJSON(r.Body)
// 		if err != nil {
// 			p.l.Println("[ERROR] deserializing product", err)
// 			http.Error(rw, "Error reading product", http.StatusBadRequest)
// 			return
// 		}

// 		// add the product to the context
// 		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
// 		r = r.WithContext(ctx)

// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(rw, r)
// 	})
// }
