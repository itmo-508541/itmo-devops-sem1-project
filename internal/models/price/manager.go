package price

import (
	"io"

	"github.com/gocarina/gocsv"
	"gopkg.in/validator.v2"
)

// csv, err := gocsv.MarshalString(prices)
// fmt.Println(csv, err)
type Manager struct {
}

func NewManager() *Manager {
	manager := &Manager{}

	return manager
}

func (m *Manager) AcceptCsv(r io.Reader) (Accepted, error) {
	accepted := Accepted{}

	prices := []PriceDTO{}
	gocsv.Unmarshal(r, &prices)
	for _, price := range prices {
		if err := validator.Validate(price); err != nil {
		} else {
		}
	}

	return accepted, nil
}
