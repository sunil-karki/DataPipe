package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"../dbconnection"

	dataStr "../interfaces"
)

// Product defines the structure for an API product
// type Product struct {
// 	Fileid      int    `json:"fileid" bson:"fileid"`
// 	Position    int    `json:"position" bson:"position"`
// 	Filename    string `json:"filename" bson:"filename"`
// 	Description string `json:"description" bson:"description"`
// 	Filedate    string `json:"filedate" bson:"filedate"`
// 	Source      string `json:"source" bson:"source"`
// }

// Product defines the structure for an API dataSource
type Product = dataStr.Product

// Products is a collection of Product
type Products []*Product

// productList is a hard coded list of products as an example of data source
var productList = []*Product{
	&dataStr.Product{
		Fileid:      1,
		Position:    1,
		Filename:    "Latte",
		Description: "Frothy milky coffee",
		Filedate:    time.Now().UTC().String(),
		Source:      "abc323",
	},
	&dataStr.Product{
		Fileid:      2,
		Position:    2,
		Filename:    "Espresso",
		Description: "Short and strong coffee without milk",
		Filedate:    time.Now().UTC().String(),
		Source:      "fjd34",
	},
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provPositiones better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// FromJSON Decodes JSON
func (p *Products) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

func GetData() Products {
	// fetch the data from the datasource
	l := log.New(os.Stdout, "DB ", log.LstdFlags)
	conn := dbconnection.NewConnection(l)
	// conn.Connect()
	// clientConn := conn.GetClient()
	// conn.InsertInterface(clientConn)
	conn.CreateConnection()
	// conn.InsertInterface()
	// conn.InsertInterface()
	// conn.UpdateInterface()
	// conn.DeleteInterface()
	return conn.ReturnRecordInterface()
	//////////////////////////////////////////////////////////////////////////////////////////////
}

// AddProduct adds product
func AddProduct(p *Product) {
	p.Position = getNextPosition()
	productList = append(productList, p)
}

// UpdateProduct updates the value of the Product
func UpdateProduct(Position int, p *Product) error {
	_, pos, err := findProduct(Position)
	if err != nil {
		return err
	}

	p.Position = Position
	productList[pos] = p

	return nil
}

// ErrProductNotFound var for "Product Not Found"
var ErrProductNotFound = fmt.Errorf("Product not found")

// findProduct to go through all Products and find a particular Product
func findProduct(Position int) (*Product, int, error) {
	for i, p := range productList {
		if p.Position == Position {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// getNextPosition returns next Position for new product
func getNextPosition() int {
	fmt.Println(len(productList))
	lp := productList[len(productList)-1]
	return lp.Position + 1
}
