package handlers

// Uploading files
// curl localhost:9090/images/1/uploaded.png -d @testpic.png
// Getting files
// curl localhost:9090/images/1/uploaded.png
// Gets/Uploads you ACII converted files

// Downloading files
// curl -v localhost:9090/images/1/testpic2.png -o file.png

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"../files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// ServeHTTP implements the http.Handler interface
func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	fmt.Println("Printing Body")
	fmt.Println(r.Body)

	// To see the content of http.Request
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error: On Ioutil ....")
	}

	fmt.Println("Printing reqbody")
	fmt.Println(string(reqbody))

	fmt.Println("-------------------------")
	fmt.Println(r.Header["Content-Type"])
	fmt.Println(r.Header["Content-Type"][0])
	fmt.Println(r.GetBody())
	fmt.Println(r.ContentLength)
	fmt.Println("-------------------------")

	io.Copy(os.Stdout, r.Body) // this line.
	fmt.Println("Printing r")
	fmt.Println()
	fmt.Println(r)

	// f.saveFile(id, fn, rw, r)
	f.saveFile(id, fn, rw, r.Body)
}

// UploadMultipart need to find its meaning
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024) // (128 * 1024) is a size Remember
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form for id", "id", id)

	if idErr != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected expected integer id", http.StatusBadRequest)
		return
	}

	// FormFile :--  Need to find its working
	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}

	fmt.Println("UploadMulti : ", ff)

	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)
}

// func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
// 	f.log.Error("Invalid path", "path", uri)
// 	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
// }

// saveFile saves the contents of the request to a file
// func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	// err := f.store.Save(fp, r.Body)
	fmt.Println()
	fmt.Println("printing r in save module")
	fmt.Println(r)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
