package utilities

import (
	"github.com/go-playground/validator"
	model "github.com/kuzeofficial/golang-test-api/models"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}