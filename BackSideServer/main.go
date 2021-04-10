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

// /////////////////////////////////////////////////////////////////////////////////////////////////
// // Help from https://github.com/Freshman-tech/file-upload
// const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

// // Progress is used to track the progress of a file upload.
// // It implements the io.Writer interface so it can be passed
// // to an io.TeeReader()
// type Progress struct {
// 	TotalSize int64
// 	BytesRead int64
// }

// // Write is used to satisfy the io.Writer interface.
// // Instead of writing somewhere, it simply aggregates
// // the total bytes on each read
// func (pr *Progress) Write(p []byte) (n int, err error) {
// 	n, err = len(p), nil
// 	pr.BytesRead += int64(n)
// 	pr.Print()
// 	return
// }

// // Print displays the current progress of the file upload
// func (pr *Progress) Print() {
// 	if pr.BytesRead == pr.TotalSize {
// 		fmt.Println("DONE!")
// 		return
// 	}

// 	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
// }

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/html")
// 	http.ServeFile(w, r, "index.html")
// }

// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Inside uploadHandler")

// 	if r.Method != "POST" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		fmt.Println("Method not allowed")
// 		return
// 	}

// 	// 32 MB is the default used by FormFile
// 	if err := r.ParseMultipartForm(32 << 20); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		fmt.Println("StatusBadRequest")
// 		return
// 	}

// 	fmt.Println("Now getting reference to fileheaders")
// 	// get a reference to the fileHeaders
// 	files := r.MultipartForm.File["file"]

// 	for _, fileHeader := range files {
// 		if fileHeader.Size > MAX_UPLOAD_SIZE {
// 			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
// 			return
// 		}

// 		file, err := fileHeader.Open()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		defer file.Close()

// 		buff := make([]byte, 512)
// 		_, err = file.Read(buff)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		filetype := http.DetectContentType(buff)
// 		if filetype != "image/jpeg" && filetype != "image/png" {
// 			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
// 			return
// 		}

// 		_, err = file.Seek(0, io.SeekStart)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		err = os.MkdirAll("./uploads", os.ModePerm)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		defer f.Close()

// 		pr := &Progress{
// 			TotalSize: fileHeader.Size,
// 		}

// 		_, err = io.Copy(f, io.TeeReader(file, pr))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	fmt.Fprintf(w, "Upload successful")
// }

//////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	l := log.New(os.Stdout, "files-api ", log.LstdFlags)

	fileupload := handlers.NewFileUploads(l)

	// mw := handlers.GzipHandler{}

	// creating a new serve mux and registering the handlers
	smux := mux.NewRouter()

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
