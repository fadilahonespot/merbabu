package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}