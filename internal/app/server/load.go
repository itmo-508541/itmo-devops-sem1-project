package server

import (
	"fmt"
	"log"
	"net/http"
	"project_sem/internal/app/report"
	"project_sem/internal/app/validate"
	"project_sem/internal/database"

	"github.com/gocarina/gocsv"
)

// NewLoadHandler возвращает GET handler
// http://localhost:8080/api/v0/prices?type=csv&start=2023-01-01&end=2025-10-01&min=10&max=20
func NewLoadHandler(conn *database.Database) http.HandlerFunc {
	reportRepo := report.NewRepository(conn)

	return func(w http.ResponseWriter, r *http.Request) {
		var csv string
		var err error

		filter := report.NewRequestFilter(r)
		fileType := r.URL.Query().Get("type")
		if fileType == "request" {
			JSONResponse(w, filter, http.StatusOK)

			return
		} else {
			v, err := validate.New()
			if err != nil {
				log.Println(fmt.Errorf("validators.NewValidate: %w", err))
				JSONInternalServerError(w)

				return
			}
			err = v.Struct(filter)
			if err != nil {
				log.Println(fmt.Errorf("v.Struct: %w", err))
				JSONBadRequestError(w)

				return
			}
		}

		all, err := reportRepo.All(r.Context(), filter)
		if err == nil {
			csv, err = gocsv.MarshalString(all)
		}
		if err != nil {
			log.Println(fmt.Errorf("ServeHTTP: %w", err))
			JSONInternalServerError(w)

			return
		}

		if fileType == "csv" {
			TextResponse(w, csv, http.StatusOK)
		} else {
			ZipResponse(w, "prices.zip", csv, "data.csv", http.StatusOK)
		}
	}
}
