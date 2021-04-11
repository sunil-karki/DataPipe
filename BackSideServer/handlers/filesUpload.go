package handlers

////////////////////////////////////////////////////////////////////////////////////////////////
// Help from https://github.com/Freshman-tech/file-upload
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"../dbconnection"
	"../env"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

var basePath = env.String("BASE_PATH", false, "./uploads", "Base path to save images")

// filesUpload will be a http.Handler
type FileUpload struct {
	l *log.Logger
}

// NewFileUploads creates a files upload handler with the given logger
func NewFileUploads(l *log.Logger) *FileUpload {
	return &FileUpload{l}
}

// Progress implements the io.Writer interface so it can be passed to an io.TeeReader()
type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/html")
// 	http.ServeFile(w, r, "index.html")
// }

func (p *FileUpload) UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("basepath : ", basePath)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		// fmt.Println(fmt.Sprintf("./%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		f, err := os.Create(fmt.Sprintf("./uploads/%s", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p.InsertFile(w, fileHeader.Filename)

		defer f.Close()
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		pr := &Progress{
			TotalSize: fileHeader.Size,
		}

		_, err = io.Copy(f, io.TeeReader(file, pr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Fprintf(w, "Upload successful")
}

// InsertFile inserts the meta data of each File uploaded to DB
func (p *FileUpload) InsertFile(w http.ResponseWriter, filename string) {
	// fetch the data from the datasource
	l := log.New(os.Stdout, "Meta", log.LstdFlags)
	fmt.Println("Filename : ", filename)
	conn := dbconnection.NewConnection(l)
	conn.InsertInterface(rand.Intn(500), rand.Intn(100), filename, "New file insertion", time.Now().Format("2006-01-02"), "File Upload")
}
