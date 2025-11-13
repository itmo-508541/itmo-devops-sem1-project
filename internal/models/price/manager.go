package price

import (
	"context"
	"io"
	"project_sem/internal/models/report"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
)

// csv, err := gocsv.MarshalString(prices)
// fmt.Println(csv, err)

type Manager struct {
	processors []PriceProcessor
	repoReport *report.Repository
}

func NewManager(r *report.Repository) *Manager {
	manager := &Manager{repoReport: r}

	return manager
}

func (m *Manager) AddProcessor(processor PriceProcessor) {
	m.processors = append(m.processors, processor)
}

func (m *Manager) AcceptCsv(ctx context.Context, r io.Reader) (AcceptedDTO, error) {
	accepted := AcceptedDTO{UUID: uuid.New()}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := make(chan PriceDTO)
	defer close(inputCh)

	outputPrices := []PriceDTO{}
	go func(outputCh chan PriceDTO) {
		for _, processor := range m.processors {
			outputCh = m.processStep(doneCh, outputCh, processor)
		}

		for price := range outputCh {
			outputPrices = append(outputPrices, price)
		}
	}(inputCh)

	inputPrices := []PriceDTO{}
	gocsv.Unmarshal(r, &inputPrices)
	for _, price := range inputPrices {
		price.GroupUUID = accepted.UUID
		price.UUID = uuid.New()
		inputCh <- price
	}

	err := m.repoReport.Accept(ctx, accepted.UUID)

	return accepted, err
}

func (m *Manager) processStep(doneCh chan struct{}, inputCh chan PriceDTO, processor PriceProcessor) chan PriceDTO {
	outputCh := make(chan PriceDTO)

	go func() {
		defer close(outputCh)

		for price := range inputCh {
			err := processor(&price)
			if err != nil {
				continue
			}

			select {
			case <-doneCh:
				return
			case outputCh <- price:
			}
		}
	}()

	return outputCh
}
