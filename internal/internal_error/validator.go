package internalerror

import "github.com/go-playground/validator/v10"

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return nil
	}
	validatorErrors := err.(validator.ValidationErrors)
	return validatorErrors
}