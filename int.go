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
	_ Validatable = (*ValidatableInt[uint])(nil)
	_ Validatable = (*ValidatableInt[uint8])(nil)
	_ Validatable = (*ValidatableInt[uint16])(nil)
	_ Validatable = (*ValidatableInt[uint32])(nil)
	_ Validatable = (*ValidatableInt[uint64])(nil)
)

type Ints interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type ValidatableInt[T Ints] struct {
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
		switch v.tag {
		case nil:
			return errors.New(fmt.Sprintf("failed validation for <%T>", v.value))
		default:
			return errors.New(fmt.Sprintf("<%s> failed validation for <%T>", *v.tag, v.value))
		}
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

func (v *ValidatableInt[T]) LT(max T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		if v.value > max {
			if len(msg) > 0 {
				return errors.New(msg[0])
			}
			switch v.tag {
			case nil:
				return errors.New(fmt.Sprintf("failed <%T> validation for <LT(%d)>", v.value, max))
			default:
				return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <LT(%d)>", *v.tag, v.value, max))
			}
		}
		return nil
	})
	return v
}

func (v *ValidatableInt[T]) GT(min T, msg ...string) *ValidatableInt[T] {
	v.rules = append(v.rules, func() error {
		if v.value < min {
			if len(msg) > 0 {
				return errors.New(msg[0])
			}
			switch v.tag {
			case nil:
				return errors.New(fmt.Sprintf("failed <%T> validation for <GT(%d)>", v.value, min))
			default:
				return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <GT(%d)>", *v.tag, v.value, min))
			}
		}
		return nil
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

func Uint() *ValidatableInt[uint] {
	return &ValidatableInt[uint]{}
}

func Uint8() *ValidatableInt[uint8] {
	return &ValidatableInt[uint8]{}
}

func Uint16() *ValidatableInt[uint16] {
	return &ValidatableInt[uint16]{}
}

func Uint32() *ValidatableInt[uint32] {
	return &ValidatableInt[uint32]{}
}

func Uint64() *ValidatableInt[uint64] {
	return &ValidatableInt[uint64]{}
}
