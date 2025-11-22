//go:build dist

package server

import "net/http"

type ServeMux struct {
	*http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{http.NewServeMux()}
}

// AddHandlerFunc adds handlerFunc with PanicRecoveryMiddleware on production
func (m *ServeMux) AddHandlerFunc(pattern string, handlerFunc http.HandlerFunc) {
	m.Handle(pattern, PanicRecoveryMiddleware(handlerFunc))
}
