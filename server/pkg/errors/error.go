package errors

import (
	"errors"

	"github.com/azacdev/go-blog/internal/providers/validation"
	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromError(err error) {
	var ValidationError validator.ValidationErrors

	if errors.As(err, &ValidationError) {
		for _, fieldError := range ValidationError {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func Add(key string, value string) {
	errorsList[key] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}
