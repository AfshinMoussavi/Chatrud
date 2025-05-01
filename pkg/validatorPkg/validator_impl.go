package validatorPkg

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type validatorImpl struct {
	validate *validator.Validate
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidator() IValidator {
	v := validator.New()
	_ = v.RegisterValidation("mobile", ValidateMobile)
	return &validatorImpl{validate: v}
}

func (v *validatorImpl) ValidateStruct(s interface{}) validator.ValidationErrors {
	err := v.validate.Struct(s)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return ve
		}
	}
	return nil
}

func (v *validatorImpl) FormatErrors(obj interface{}, errs validator.ValidationErrors) []ErrorResponse {
	var response []ErrorResponse
	for _, e := range errs {
		field := getJSONFieldName(obj, e.StructField())
		msg := translateTag(field, e)
		response = append(response, ErrorResponse{
			Field:   field,
			Message: msg,
		})
	}
	return response
}

func getJSONFieldName(obj interface{}, field string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if f, ok := t.FieldByName(field); ok {
		return f.Tag.Get("json")
	}
	return field
}
