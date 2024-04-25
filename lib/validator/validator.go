package validator

import (
	"errors"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type ValidationError struct {
	Message string                 `json:"error" xml:"Error"`
	Details ValidationErrorDetails `json:"details" xml:"Details"`
}

const (
	RequiredTag string = "required"
	PhoneTag    string = "e164"
)

func (v ValidationError) Error() string {
	return v.Message
}

type ValidationErrorDetails struct {
	MissingProperties []string          `json:"missing_properties" xml:"MissingProperties"`
	InvalidProperties map[string]string `json:"invalid_properties" xml:"InvalidProperties"`
}

var (
	Validator = &DefaultValidator{}

	// compile-time interface-fulfilment check.
	_ binding.StructValidator = &DefaultValidator{}
)

// ValidateStruct always returns a ValidationError or nil.
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	v.lazyinit()

	if err := v.validate.Struct(obj); err != nil {
		if e := (&validator.InvalidValidationError{}); errors.As(err, &e) {
			return err
		}

		validationErr := ValidationError{
			Message: "invalid_body",
			Details: ValidationErrorDetails{
				MissingProperties: make([]string, 0),
				InvalidProperties: make(map[string]string),
			},
		}
		for _, err := range err.(validator.ValidationErrors) {
			name := RemovePrefix(err.StructNamespace())
			switch err.Tag() {
			case RequiredTag:
				validationErr.Details.MissingProperties = append(validationErr.Details.MissingProperties, name)
			case PhoneTag:
				validationErr.Details.InvalidProperties[name] = "phone"
			default:
				validationErr.Details.InvalidProperties[name] = err.Tag()
			}

		}

		return validationErr
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("validate")

		// add any custom validations etc. here
	})
}

func RemovePrefix(namespace string) string {
	index := strings.Index(namespace, ".") + 1
	return namespace[index:]
}
