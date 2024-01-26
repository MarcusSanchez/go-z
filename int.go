package z

import (
	"errors"
	"fmt"
	"strings"
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

type ValidatableInt[T ints] struct {
	tag    *string
	value  T
	rules  []Rule
	errors []string
}

func (v *ValidatableInt[T]) Validate(data any, tags ...string) error {
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

func (v *ValidatableInt[T]) Lt(max T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value < max:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Lt(%d)>", *v.tag, v.value, max))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Lt(%d)>", v.value, max))
		}
	})
	return v
}

func (v *ValidatableInt[T]) Gt(min T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value > min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Gt(%d)>", *v.tag, v.value, min))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Gt(%d)>", v.value, min))
		}
	})
	return v
}

func (v *ValidatableInt[T]) Lte(max T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value <= max:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Lte(%d)>", *v.tag, v.value, max))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Lte(%d)>", v.value, max))
		}
	})
	return v
}

func (v *ValidatableInt[T]) Gte(min T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value >= min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Gte(%d)>", *v.tag, v.value, min))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Gte(%d)>", v.value, min))
		}
	})
	return v
}

func (v *ValidatableInt[T]) Eq(to T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value == to:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Eq(%d)>", *v.tag, v.value, to))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Eq(%d)>", v.value, to))
		}
	})
	return v
}

func (v *ValidatableInt[T]) NotEq(to T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value != to:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <NotEq(%d)>", *v.tag, v.value, to))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <NotEq(%d)>", v.value, to))
		}
	})
	return v
}

func (v *ValidatableInt[T]) Positive(msg ...string) *ValidatableInt[T] {
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

func (v *ValidatableInt[T]) Negative(msg ...string) *ValidatableInt[T] {
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

func (v *ValidatableInt[T]) NonNegative(msg ...string) *ValidatableInt[T] {
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

func (v *ValidatableInt[T]) NonPositive(msg ...string) *ValidatableInt[T] {
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

func (v *ValidatableInt[T]) NonZero(msg ...string) *ValidatableInt[T] {
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

func (v *ValidatableInt[T]) Custom(rule func(T) bool, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case rule(v.value):
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Custom>", *v.tag, v.value))
		default:
			return errors.New(fmt.Sprintf("failed <%T> validation for <Custom>", v.value))
		}
	})
	return v
}

func Int() *ValidatableInt[int] {
	return &ValidatableInt[int]{}
}

func Int8() *ValidatableInt[int8] {
	return &ValidatableInt[int8]{}
}

func Int16() *ValidatableInt[int16] {
	return &ValidatableInt[int16]{}
}

func Int32() *ValidatableInt[int32] {
	return &ValidatableInt[int32]{}
}

func Int64() *ValidatableInt[int64] {
	return &ValidatableInt[int64]{}
}
