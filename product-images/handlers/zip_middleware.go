package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWarpedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}
		next.ServeHTTP(rw, r)
	})
}

type WarpedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWarpedResponseWriter(rw http.ResponseWriter) *WarpedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WarpedResponseWriter{rw: rw, gw: gw}
}

func (wr *WarpedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WarpedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WarpedResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}

func (wr *WarpedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
