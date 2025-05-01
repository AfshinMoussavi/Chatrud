package validatorPkg

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func translateTag(field string, e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must have at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", field, e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", field)
	case "numeric":
		return fmt.Sprintf("%s must contain only numeric characters", field)
	case "mobile":
		return fmt.Sprintf("%s must be a valid Iranian phone number starting with 09", field)
	case "unique":
		return fmt.Sprintf("%s must be unique", field)
	default:
		return fmt.Sprintf("Invalid value for %s", field)
	}
}
