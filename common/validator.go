package common

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) (bool, error) {
	err := validate.Struct(s)
	if err != nil {
		return false, err
	}
	return true, nil
}
