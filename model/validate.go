package model

import "github.com/go-playground/validator/v10"


func ValidateStruct[T any](payload T) []*ValidationError {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		var errors []*ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
		return errors
	}
	return nil
}

type ValidationError struct {
	Field string
	Tag         string
	Value       string
}
