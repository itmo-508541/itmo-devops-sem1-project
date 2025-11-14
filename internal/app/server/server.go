package server

import (
	"net/http"
	"time"
)

func NewWebServer(mux *http.ServeMux, addr string) *http.Server {
	return &http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
