//go:build !dist

package server

import (
	"net/http"
	"project_sem/internal/app/assets"
)

type ServeMux struct {
	*http.ServeMux
}

func NewServeMux() *ServeMux {
	mux := &ServeMux{http.NewServeMux()}

	mux.Handle("GET /favicon.ico", http.FileServer(http.FS(assets.FaviconFS)))
	mux.Handle("GET /", http.FileServer(http.FS(assets.IndexFS)))

	return mux
}

func (m *ServeMux) AddHandlerFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Handle(pattern, handlerFunc)
}
