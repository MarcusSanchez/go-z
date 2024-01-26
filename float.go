package z

import (
	"errors"
	"fmt"
	"strings"
)

var (
	_ Validatable = (*ValidatableFloat[float32])(nil)
	_ Validatable = (*ValidatableFloat[float64])(nil)
)

type floats interface {
	float32 | float64
}

type ValidatableFloat[T floats] struct {
	tag    *string
	value  T
	rules  []Rule
	errors []string
}

func (v *ValidatableFloat[T]) Validate(data any, tags ...string) error {
	if len(tags) > 0 {
		v.tag = &tags[0]
	}
	var ok bool
	if v.value, ok = data.(T); !ok {
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed validation for <%T>", v.value))
		}
		return errors.New(fmt.Sprintf("<%s> failed validation for <%T>", *v.tag, v.value))
	}
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			v.errors = append(v.errors, err.Error())
		}
	}
	if len(v.errors) > 0 {
		return errors.New(strings.Join(v.errors, ", "))
	}
	return nil
}

func (v *ValidatableFloat[T]) Lt(max T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value < max:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Lt(%g)>", *v.tag, v.value, max))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Lt(%g)>", v.value, max))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Gt(min T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value > min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Gt(%g)>", *v.tag, v.value, min))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Gt(%g)>", v.value, min))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Lte(max T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value <= max:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Lte(%g)>", *v.tag, v.value, max))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Lte(%g)>", v.value, max))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Gte(min T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value >= min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Gte(%g)>", *v.tag, v.value, min))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Gte(%g)>", v.value, min))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Eq(to T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value == to:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Eq(%g)>", *v.tag, v.value, to))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Eq(%g)>", v.value, to))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) NotEq(to T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value != to:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <NotEq(%g)>", *v.tag, v.value, to))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <NotEq(%g)>", v.value, to))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Positive(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value > 0:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Positive>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Positive>", v.value))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) Negative(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value < 0:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Negative>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Negative>", v.value))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) NonNegative(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value >= 0:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <NonNegative>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <NonNegative>", v.value))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) NonPositive(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value <= 0:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <NonPositive>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <NonPositive>", v.value))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) NonZero(msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value != 0:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <NonZero>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <NonZero>", v.value))
		}
	})
	return v
}

func Float64() *ValidatableFloat[float64] {
	return &ValidatableFloat[float64]{}
}

func Float32() *ValidatableFloat[float32] {
	return &ValidatableFloat[float32]{}
}
