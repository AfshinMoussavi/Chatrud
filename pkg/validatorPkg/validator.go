package validatorPkg

import "github.com/go-playground/validator/v10"

type IValidator interface {
	ValidateStruct(s interface{}) validator.ValidationErrors
	FormatErrors(obj interface{}, errs validator.ValidationErrors) []ErrorResponse
}
