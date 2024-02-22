package z

import (
	"fmt"
	"github.com/MarcusSanchez/go-z/internal"
)

var (
	_ Validatable = (*ValidatableUint[uint])(nil)
	_ Validatable = (*ValidatableUint[uint8])(nil)
	_ Validatable = (*ValidatableUint[uint16])(nil)
	_ Validatable = (*ValidatableUint[uint32])(nil)
	_ Validatable = (*ValidatableUint[uint64])(nil)
)

type uints interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// ValidatableUint is a uint, uint8, uint16, uint32, or uint64 that can be validated.
type ValidatableUint[T uints] struct {
	tag      *string
	value    T
	rules    []rule
	optional bool
}

// Validate validates a uint, uint8, uint16, uint32, or uint64 against its schema.
//
//	Returns Errors if:
//	=> data is not a uint, uint8, uint16, uint32, or uint64 or doesn't match the correct generic type
//	=> data fails any of the schema's rules
func (v *ValidatableUint[T]) Validate(data any, tag ...string) Errors {
	if len(tag) > 0 {
		v.tag = &tag[0]
	}
	if v.optional {
		if data == nil {
			return nil
		}
		if value, ok := data.(*T); ok {
			if value == nil {
				return nil
			}
			data = *value
		}
	}
	var ok bool
	if v.value, ok = data.(T); !ok {
		if v.tag == nil {
			return internal.NewValidationErrors(fmt.Sprintf("failed validation for <%T>", v.value))
		}
		return internal.NewValidationErrors(fmt.Sprintf("<%s> failed validation for <%T>", *v.tag, v.value))
	}
	vErrors := make([]string, 0, len(v.rules))
	for _, rule := range v.rules {
		if err := rule(); err != "" {
			vErrors = append(vErrors, err)
		}
	}
	if len(vErrors) > 0 {
		return internal.NewValidationErrors(vErrors...)
	}
	return nil
}

// Optional marks the uint as optional. Calling Validate with nil or a nil uint pointer will skip validation.
func (v *ValidatableUint[T]) Optional() *ValidatableUint[T] {
	v.optional = true
	return v
}

// Lt appends a rule validating that data is less than the provided max. (data < max)
func (v *ValidatableUint[T]) Lt(max T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value < max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Lt(%d)>", *v.tag, v.value, max)
		default:
			return fmt.Sprintf("failed <%T> validation for <Lt(%d)>", v.value, max)
		}
	})
	return v
}

// Gt appends a rule validating that data is greater than the provided min. (data > min)
func (v *ValidatableUint[T]) Gt(min T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value > min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Gt(%d)>", *v.tag, v.value, min)
		default:
			return fmt.Sprintf("failed <%T> validation for <Gt(%d)>", v.value, min)
		}
	})
	return v
}

// Lte appends a rule validating that data is less than or equal to the provided max. (data <= max)
func (v *ValidatableUint[T]) Lte(max T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value <= max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Lte(%d)>", *v.tag, v.value, max)
		default:
			return fmt.Sprintf("failed <%T> validation for <Lte(%d)>", v.value, max)
		}
	})
	return v
}

// Gte appends a rule validating that data is greater than or equal to the provided min. (data >= min)
func (v *ValidatableUint[T]) Gte(min T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value >= min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Gte(%d)>", *v.tag, v.value, min)
		default:
			return fmt.Sprintf("<%T> failed validation for <Gte(%d)>", v.value, min)
		}
	})
	return v
}

// Range appends a rule validating that data is within the provided range. (min <= data <= max)
func (v *ValidatableUint[T]) Range(min, max T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case min <= v.value && v.value <= max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Range(%d, %d)>", *v.tag, v.value, min, max)
		default:
			return fmt.Sprintf("failed <%T> validation for <Range(%d, %d)>", v.value, min, max)
		}
	})
	return v
}

// Eq appends a rule validating that data is equal to the provided value. (data == to)
func (v *ValidatableUint[T]) Eq(to T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value == to:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Eq(%d)>", *v.tag, v.value, to)
		default:
			return fmt.Sprintf("failed <%T> validation for <Eq(%d)>", v.value, to)
		}
	})
	return v
}

// NotEq appends a rule validating that data is not equal to the provided value. (data != to)
func (v *ValidatableUint[T]) NotEq(to T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value != to:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <NotEq(%d)>", *v.tag, v.value, to)
		default:
			return fmt.Sprintf("failed <%T> validation for <NotEq(%d)>", v.value, to)
		}
	})
	return v
}

// NonZero appends a rule validating that data is not equal to zero. (data != 0)
func (v *ValidatableUint[T]) NonZero(msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value != 0:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <NonZero>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <NonZero>", v.value)
		}
	})
	return v
}

// In appends a rule validating that data is in the provided slice of values.
func (v *ValidatableUint[T]) In(values []T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		for _, value := range values {
			if v.value == value {
				return ""
			}
		}
		if len(msg) > 0 {
			return msg[0]
		}
		if v.tag != nil {
			return fmt.Sprintf("<%s> failed <%T> validation for <In(%s)>", *v.tag, v.value, values)
		}
		return fmt.Sprintf("failed <%T> validation for <In(%s)>", v.value, values)
	})
	return v
}

// Custom appends a custom rule to the schema. Validates if the provided function returns true when passed data.
func (v *ValidatableUint[T]) Custom(rule func(T) bool, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case rule(v.value):
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Custom>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <Custom>", v.value)
		}
	})
	return v
}

// Uint returns a ValidatableUint[uint] for validating a uint.
func Uint() *ValidatableUint[uint] { return &ValidatableUint[uint]{} }

// Uint8 returns a ValidatableUint[uint8] for validating a uint8.
func Uint8() *ValidatableUint[uint8] { return &ValidatableUint[uint8]{} }

// Uint16 returns a ValidatableUint[uint16] for validating a uint16.
func Uint16() *ValidatableUint[uint16] { return &ValidatableUint[uint16]{} }

// Uint32 returns a ValidatableUint[uint32] for validating a uint32.
func Uint32() *ValidatableUint[uint32] { return &ValidatableUint[uint32]{} }

// Uint64 returns a ValidatableUint[uint64] for validating a uint64.
func Uint64() *ValidatableUint[uint64] { return &ValidatableUint[uint64]{} }
