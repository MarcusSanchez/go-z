package z

import (
	"errors"
	"fmt"
	"strings"
)

type Floats interface {
	float32 | float64
}

type ValidatableFloat[T Floats] struct {
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

func (v *ValidatableFloat[T]) LT(max T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		if !(v.value > max) {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		switch v.tag {
		case nil:
			return errors.New(fmt.Sprintf("failed <%T> validation for <LT(%g)>", v.value, max))
		default:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <LT(%g)>", *v.tag, v.value, max))
		}
	})
	return v
}

func (v *ValidatableFloat[T]) GT(min T, msg ...string) *ValidatableFloat[T] {
	v.rules = append(v.rules, func() error {
		if !(v.value < min) {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		switch v.tag {
		case nil:
			return errors.New(fmt.Sprintf("failed <%T> validation for <GT(%g)>", v.value, min))
		default:
			return errors.New(fmt.Sprintf("<%s> failed <%T> validation for <GT(%g)>", *v.tag, v.value, min))
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
