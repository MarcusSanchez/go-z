package internal

import (
	"strings"
)

type ValidationErrors struct {
	Errors []string
}

func NewValidationErrors(errs ...string) *ValidationErrors {
	return &ValidationErrors{Errors: errs}
}

func (e *ValidationErrors) One() string {
	if len(e.Errors) == 0 {
		panic("no validation Errors to return")
	}
	return e.Errors[0]
}

func (e *ValidationErrors) All() []string {
	if len(e.Errors) == 0 {
		panic("no validation Errors to return")
	}
	return e.Errors
}

func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		panic("no validation Errors to return")
	}
	return strings.Join(e.Errors, " | ")
}
