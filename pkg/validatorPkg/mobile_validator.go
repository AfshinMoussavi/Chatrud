package validatorPkg

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	regex := regexp.MustCompile(`^09[0-9]{9}$`)
	return regex.MatchString(phone)
}
