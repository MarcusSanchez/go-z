// Package z, inspired by zod, is a library for validating structs and other primitives.
package z

type Validatable interface {
	Validate(data any, tags ...string) error
}

type Rule func() error
