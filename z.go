// Package z, inspired by zod, is a library for validating structs and other primitives.
package z

import "strings"

type Validatable interface {
	Validate(data any, tags ...string) error
}

type Rule func() error

func ErrSlice(err error) (errs []string) {
	vErrs := strings.Split(err.Error(), ", ")
	for _, e := range vErrs {
		errs = append(errs, e)
	}
	return errs
}
