// Package z, inspired by zod, is a library for validating structs and other primitives.
package z

import "github.com/MarcusSanchez/go-z/internal"

// Validatable interface is implemented by all z-primitives.
type Validatable interface {
	// Validate validates a primitive against its schema. Returns Errors if any rules fail.
	Validate(data any, tags ...string) Errors
}

// Errors interface is z's custom error type. It is returned by all of z-primitives' Validate methods.
type Errors interface {
	// One returns the first failed validation message. If schema is a struct, it will return the
	// first failed validation message of the tag with alphabetical priority (a->z).
	One() string
	// All returns all failed validation messages in a slice.
	All() []string
	// Error is for compatibility with the error interface.
	// It returns a string representation of All() joined by pipes.
	Error() string
}

type rule func() string

var _ Errors = (*internal.ValidationErrors)(nil)
