package validate

import "github.com/go-playground/validator/v10"

func New() (*validator.Validate, error) {
	v := validator.New()
	if err := v.RegisterValidation("date", dateValidator()); err != nil {
		return nil, err
	}
	if err := v.RegisterValidation("notblank", notBlankValidator()); err != nil {
		return nil, err
	}

	return v, nil
}
