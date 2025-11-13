package price

import (
	"project_sem/internal/models/report"

	"github.com/google/uuid"
)

type PriceDTO struct {
	report.ReportDTO
	GroupUUID uuid.UUID
}

type AcceptedDTO struct {
	UUID uuid.UUID `json:"uuid"`
}
