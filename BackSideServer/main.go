package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./env"
	"./handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

////////// Section For file uploading and serving ///////////////////////////////////////////////
var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

/////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	l := log.New(os.Stdout, "files-api ", log.LstdFlags)

	fileupload := handlers.NewFileUploads(l)
	ph := handlers.NewProducts(l)

	// mw := handlers.GzipHandler{}

	// creating a new serve mux and registering the handlers
	smux := mux.NewRouter()
	//////////////////////////////////////////////////////////////////////////////////////////////

	getRouter := smux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetData)

	putRouter := smux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	// putRouter.Use(ph.MiddlewareValidateProduct)
	// Middleware implements first then the UpdateProducts handler starts to work.

	postRouter := smux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	// postRouter.Use(ph.MiddlewareValidateProduct)

	///////// Section For file uploading and serving//////////////////////////////////////////////
	phf := smux.Methods(http.MethodPost).Subrouter()
	phf.HandleFunc("/upload", fileupload.UploadHandler)
	//////////////////////////////////////////////////////////////////////////////////////////////

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(smux)

	s := &http.Server{
		Addr:    ":9090",
		Handler: handler,
		// ErrorLog:     sl,                // the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	l.Println("BackServer RUNNING ...")

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println()
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// The goroutine that WithCancel or WithTimeout created will be retained in memory indefinitely (until the program shuts down), causing a memory leak.
	// If done repeatedly, memory will balloon significantly. So defer cancel() done.
	defer cancel()

	s.Shutdown(tc)

}
