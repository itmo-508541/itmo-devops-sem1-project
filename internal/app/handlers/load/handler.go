package load

import (
	"net/http"
	"project_sem/internal/server"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.JSONResponse(w, outDTO{Count: 5}, http.StatusOK)
}
