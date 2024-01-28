package z

import (
	"fmt"
	"github.com/MarcusSanchez/go-z/internal"
)

var (
	_ Validatable = (*ValidatableFloat[float32])(nil)
	_ Validatable = (*ValidatableFloat[float64])(nil)
)

type floats interface {
	float32 | float64
}

// ValidatableFloat is a float32 or float64 that can be validated.
type ValidatableFloat[T floats] struct {
	tag      *string
	value    T
	rules    []rule
	optional bool
}

// Validate validates a float32 or float64 against its schema.
//
//	Returns Errors if:
//	=> data is not a float32/float64 or doesn't match the correct generic type
//	=> data fails any of the schema's rules
func (v *ValidatableFloat[T]) Validate(data any, tag ...string) Errors {
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

// Optional marks the float32 or float64 as optional. Calling Validate with nil or a nil float32/float64 pointer will skip validation.
func (v *ValidatableFloat[T]) Optional() *ValidatableFloat[T] {
	v.optional = true
	return v
}

// Lt appends a rule validating that data is less than the provided max. (data < max)
func (v *ValidatableFloat[T]) Lt(max T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value < max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Lt(%g)>", *v.tag, v.value, max)
		default:
			return fmt.Sprintf("failed <%T> validation for <Lt(%g)>", v.value, max)
		}
	})
	return v
}

// Gt appends a rule validating that data is greater than the provided min. (data > min)
func (v *ValidatableFloat[T]) Gt(min T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value > min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Gt(%g)>", *v.tag, v.value, min)
		default:
			return fmt.Sprintf("failed <%T> validation for <Gt(%g)>", v.value, min)
		}
	})
	return v
}

// Lte appends a rule validating that data is less than or equal to the provided max. (data <= max)
func (v *ValidatableFloat[T]) Lte(max T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value <= max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Lte(%g)>", *v.tag, v.value, max)
		default:
			return fmt.Sprintf("failed <%T> validation for <Lte(%g)>", v.value, max)
		}
	})
	return v
}

// Gte appends a rule validating that data is greater than or equal to the provided min. (data >= min)
func (v *ValidatableFloat[T]) Gte(min T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value >= min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Gte(%g)>", *v.tag, v.value, min)
		default:
			return fmt.Sprintf("failed <%T> validation for <Gte(%g)>", v.value, min)
		}
	})
	return v
}

// Eq appends a rule validating that data is equal to the provided value. (data == to)
func (v *ValidatableFloat[T]) Eq(to T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value == to:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Eq(%g)>", *v.tag, v.value, to)
		default:
			return fmt.Sprintf("failed <%T> validation for <Eq(%g)>", v.value, to)
		}
	})
	return v
}

// NotEq appends a rule validating that data is not equal to the provided value. (data != to)
func (v *ValidatableFloat[T]) NotEq(to T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value != to:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <NotEq(%g)>", *v.tag, v.value, to)
		default:
			return fmt.Sprintf("failed <%T> validation for <NotEq(%g)>", v.value, to)
		}
	})
	return v
}

// Positive appends a rule validating that data is greater than zero. (data > 0)
func (v *ValidatableFloat[T]) Positive(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value > 0:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Positive>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <Positive>", v.value)
		}
	})
	return v
}

// Negative appends a rule validating that data is less than zero. (data < 0)
func (v *ValidatableFloat[T]) Negative(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value < 0:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Negative>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <Negative>", v.value)
		}
	})
	return v
}

// NonNegative appends a rule validating that data is greater than or equal to zero. (data >= 0)
func (v *ValidatableFloat[T]) NonNegative(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value >= 0:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <NonNegative>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <NonNegative>", v.value)
		}
	})
	return v
}

// NonPositive appends a rule validating that data is less than or equal to zero. (data <= 0)
func (v *ValidatableFloat[T]) NonPositive(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value <= 0:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <NonPositive>", *v.tag, v.value)
		default:
			return fmt.Sprintf("failed <%T> validation for <NonPositive>", v.value)
		}
	})
	return v
}

// NonZero appends a rule validating that data is not equal to zero. (data != 0)
func (v *ValidatableFloat[T]) NonZero(msg ...string) *ValidatableFloat[T] {
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

// Custom appends a custom rule to the schema. Validates if the provided function returns true when passed data.
func (v *ValidatableFloat[T]) Custom(rule func(T) bool, msg ...string) *ValidatableFloat[T] {
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

// Float32 returns a ValidatableFloat[float32] for validating an float32.
func Float32() *ValidatableFloat[float32] { return &ValidatableFloat[float32]{} }

// Float64 returns a ValidatableFloat[float64] for validating an float64.
func Float64() *ValidatableFloat[float64] { return &ValidatableFloat[float64]{} }
