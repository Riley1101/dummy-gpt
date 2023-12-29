package common

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return false, err
	}
	return true, nil
}
