package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
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
	// http.HandleFunc("/", handler)
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewAbout(l)
	ph := handlers.NewProducts(l)

	smux := http.NewServeMux()
	smux.Handle("/", ph)
	// smux.Handle("/products", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      smux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// http.ListenAndServe(":9090", nil)
	// http.ListenAndServe(":9090", smux)
	// s.ListenAndServe()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// The goroutine that WithCancel or WithTimeout created will be retained in memory indefinitely (until the program shuts down), causing a memory leak.
	// If done repeatedly, memory will balloon significantly. So defer cancel() done.
	defer cancel()

	s.Shutdown(tc)

}
