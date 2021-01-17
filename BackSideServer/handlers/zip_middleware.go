package handlers

// Copied from https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_12/product-images/handlers/zip_middleware.go
// Author Nic Jackson

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GzipHandler struct
type GzipHandler struct {
}

// GzipMiddleware handles gzip
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gziped response
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()

			return
		}

		// handle normal
		next.ServeHTTP(rw, r)
	})
}

// WrappedReponseWriter struct
type WrappedReponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter returns WrappedReponseWriter
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedReponseWriter {
	gw := gzip.NewWriter(rw)

	return &WrappedReponseWriter{rw: rw, gw: gw}
}

// Header function
func (wr *WrappedReponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedReponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader function
func (wr *WrappedReponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

// Flush uses Gzip Flush
func (wr *WrappedReponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
