package server

type errorResponseDTO struct {
	Error errorDTO `json:"error"`
}

type errorDTO struct {
	Message string `json:"message"`
}
