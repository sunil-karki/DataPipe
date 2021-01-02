package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		log.Printf("Logging: Found Error in Reading...")
		fmt.Fprintf(os.Stdout, "Responding: Found Error in  Reading...\n")

	} else {
		log.Printf("Logging: Data %s", d)
		fmt.Fprintf(os.Stdout, "Responding to Server: Data %s\n", d)
		fmt.Fprintf(w, "Responding to User: Data %s\n", d)
	}

}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":9090", nil)
}
