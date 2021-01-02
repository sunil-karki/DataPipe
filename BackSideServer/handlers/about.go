package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// About is Struct for About
type About struct {
	l *log.Logger
}

// NewAbout func : Need to find its working
func NewAbout(l *log.Logger) *About {
	return &About{l}
}

func (h *About) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inside Handler for %s!\n", r.URL.Path[1:])
	h.l.Println("New Log Test" + " : Works !!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
		log.Printf("About --- Logging: Found Error in Reading...")
		fmt.Fprintf(os.Stdout, "About --- Responding: Found Error in  Reading...\n")
		return

	}
	// else {
	log.Printf("About --- Logging: Data %s", d)
	fmt.Fprintf(os.Stdout, "About --- Responding to Server: Data %s\n", d)
	fmt.Fprintf(w, "About --- Responding to User: Data %s\n", d)
	// return
	// }
	// return
}
