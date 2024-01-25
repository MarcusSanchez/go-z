package z

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var _ Validatable = (*ValidatableString)(nil)

type ValidatableString struct {
	tag    *string
	value  string
	rules  []Rule
	errors []string
}

func (v *ValidatableString) Validate(data any, tag ...string) error {
	if len(tag) > 0 {
		v.tag = &tag[0]
	}
	var ok bool
	if v.value, ok = data.(string); !ok {
		switch v.tag {
		case nil:
			return errors.New("failed validation for <String>")
		default:
			return errors.New(fmt.Sprintf("<%s> failed validation for <String>", *v.tag))
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

func (v *ValidatableString) Email(msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
		if !emailRegex.MatchString(v.value) {
			if len(msg) > 0 {
				return errors.New(msg[0])
			}
			switch v.tag {
			case nil:
				return errors.New("failed <String> validation for <Email>")
			default:
				return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Email>", *v.tag))
			}
		}
		return nil
	})
	return v
}

func (v *ValidatableString) Min(min int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		if len(v.value) < min {
			if len(msg) > 0 {
				return errors.New(msg[0])
			}
			switch v.tag {
			case nil:
				return errors.New(fmt.Sprintf("failed <String> validation for <LT(%d)>", min))
			default:
				return errors.New(fmt.Sprintf("<%s> failed <String> validation for <LT(%d)>", *v.tag, min))
			}
		}
		return nil
	})
	return v
}

func (v *ValidatableString) Max(max int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		if len(v.value) > max {
			if len(msg) > 0 {
				return errors.New(msg[0])
			}
			switch v.tag {
			case nil:
				return errors.New(fmt.Sprintf("failed <String> validation for <GT(%d)>", max))
			default:
				return errors.New(fmt.Sprintf("<%s> failed <String> validation for <GT(%d)>", *v.tag, max))
			}
		}
		return nil
	})
	return v
}

func String() *ValidatableString {
	return &ValidatableString{}
}
