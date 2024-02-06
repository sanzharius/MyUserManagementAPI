package controller

import (
	"github.com/go-playground/validator/v10"
	"myAPIProject/internal/apperrors"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return apperrors.ValidatorCustomValidatorValidateErr.AppendMessage(err)
	}

	return nil
}
