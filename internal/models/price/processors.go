package price

import (
	"context"

	"gopkg.in/validator.v2"
)

type PriceProcessor func(price *PriceDTO) error

func NewValidateProcessor() PriceProcessor {
	return func(price *PriceDTO) error {
		return validator.Validate(price)
	}
}

// @todo заменить на интерфейс, чтобы можно было протестировать.
func NewPersistProcessor(ctx context.Context, repo *Repository) PriceProcessor {
	return func(price *PriceDTO) error {
		return repo.Insert(ctx, price)
	}
}
