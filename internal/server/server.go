package server

import (
	"net/http"
	"time"
)

func New(mux *http.ServeMux, addr string) *http.Server {
	// @todo тут куда-то нужно добавить контекст general:context, наверное?

	return &http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
