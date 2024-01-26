package z

import (
	"errors"
	"fmt"
	"strings"
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

type ValidatableUint[T uints] struct {
	tag    *string
	value  T
	rules  []Rule
	errors []string
}

func (v *ValidatableUint[T]) Validate(data any, tags ...string) error {
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

func (v *ValidatableUint[T]) Lt(max T, msg ...string) *ValidatableUint[T] {
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

func (v *ValidatableUint[T]) Gt(min T, msg ...string) *ValidatableUint[T] {
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

func (v *ValidatableUint[T]) Lte(max T, msg ...string) *ValidatableUint[T] {
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

func (v *ValidatableUint[T]) Gte(min T, msg ...string) *ValidatableUint[T] {
	v.rules = append(v.rules, func() error {
		switch {
		case v.value >= min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <Gte(%d)>", *v.tag, v.value, min))
		default:
			return errors.New(fmt.Sprintf("<%T> failed validation for <Gte(%d)>", v.value, min))
		}
	})
	return v
}

func (v *ValidatableUint[T]) Eq(to T, msg ...string) *ValidatableUint[T] {
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

func (v *ValidatableUint[T]) NotEq(to T, msg ...string) *ValidatableUint[T] {
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

func (v *ValidatableUint[T]) NonZero(msg ...string) *ValidatableUint[T] {
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

func Uint() *ValidatableUint[uint] {
	return &ValidatableUint[uint]{}
}

func Uint8() *ValidatableUint[uint8] {
	return &ValidatableUint[uint8]{}
}

func Uint16() *ValidatableUint[uint16] {
	return &ValidatableUint[uint16]{}
}

func Uint32() *ValidatableUint[uint32] {
	return &ValidatableUint[uint32]{}
}

func Uint64() *ValidatableUint[uint64] {
	return &ValidatableUint[uint64]{}
}
