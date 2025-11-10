package load

import "project_sem/internal/server"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) HandlerFunc() server.Handler[inDTO, outDTO] {
	return func(in inDTO, out *outDTO) error {
		out.Count = in.Count

		return nil
	}
}
