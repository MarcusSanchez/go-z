package z

import (
	"fmt"
	"github.com/MarcusSanchez/go-z/internal"
)

var (
	_ Validatable = (*ValidatableInt[int])(nil)
	_ Validatable = (*ValidatableInt[int8])(nil)
	_ Validatable = (*ValidatableInt[int16])(nil)
	_ Validatable = (*ValidatableInt[int32])(nil)
	_ Validatable = (*ValidatableInt[int64])(nil)
)

type ints interface {
	int | int8 | int16 | int32 | int64
}

// ValidatableInt is an int, int8, int16, int32, or int64 that can be validated.
type ValidatableInt[T ints] struct {
	tag      *string
	value    T
	rules    []rule
	optional bool
}

// Validate validates an int, int8, int16, int32, or int64 against its schema.
//
//	Returns Errors if:
//	=> data is not an int, int8, int16, int32, or int64 or doesn't match the correct generic type
//	=> data fails any of the schema's rules
func (v *ValidatableInt[T]) Validate(data any, tag ...string) Errors {
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

// Optional marks the int as optional. Calling Validate with nil or a nil int pointer will skip validation.
func (v *ValidatableInt[T]) Optional() *ValidatableInt[T] {
	v.optional = true
	return v
}

// Lt appends a rule validating that data is less than the provided max. (data < max)
func (v *ValidatableInt[T]) Lt(max T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Gt(min T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Lte(max T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Gte(min T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() string {
		switch {
		case v.value >= min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <%T> validation for <Gte(%d)>", *v.tag, v.value, min)
		default:
			return fmt.Sprintf("failed <%T> validation for <Gte(%d)>", v.value, min)
		}
	})
	return v
}

// Range appends a rule validating that data is within the provided range. (min <= data <= max)
func (v *ValidatableInt[T]) Range(min, max T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Eq(to T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) NotEq(to T, msg ...string) *ValidatableInt[T] {
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

// Positive appends a rule validating that data is greater than zero. (data > 0)
func (v *ValidatableInt[T]) Positive(msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Negative(msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) NonNegative(msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) NonPositive(msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) NonZero(msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) In(values []T, msg ...string) *ValidatableInt[T] {
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
func (v *ValidatableInt[T]) Custom(rule func(T) bool, msg ...string) *ValidatableInt[T] {
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

// Int returns a ValidatableInt[int] for validating an int.
func Int() *ValidatableInt[int] { return &ValidatableInt[int]{} }

// Int8 returns a ValidatableInt[int8] for validating an int8.
func Int8() *ValidatableInt[int8] { return &ValidatableInt[int8]{} }

// Int16 returns a ValidatableInt[int16] for validating an int16.
func Int16() *ValidatableInt[int16] { return &ValidatableInt[int16]{} }

// Int32 returns a ValidatableInt[int32] for validating an int32.
func Int32() *ValidatableInt[int32] { return &ValidatableInt[int32]{} }

// Int64 returns a ValidatableInt[int64] for validating an int64.
func Int64() *ValidatableInt[int64] { return &ValidatableInt[int64]{} }
