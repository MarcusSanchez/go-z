package z

import (
	"errors"
	"fmt"
	"net/mail"
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
		if v.tag == nil {
			return errors.New("failed validation for <String>")
		}
		return errors.New(fmt.Sprintf("<%s> failed validation for <String>", *v.tag))
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

func (v *ValidatableString) Min(min int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		switch {
		case len(v.value) > min:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Min(%d)>", *v.tag, min))
		default:
			return errors.New(fmt.Sprintf("failed <String> validation for <Min(%d)>", min))
		}
	})
	return v
}

func (v *ValidatableString) Max(max int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		switch {
		case len(v.value) < max:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Max(%d)>", *v.tag, max))
		default:
			return errors.New(fmt.Sprintf("failed <String> validation for <Max(%d)>", max))
		}
	})
	return v
}

func (v *ValidatableString) Email(msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		a, err := mail.ParseAddress(v.value)
		switch {
		case err == nil && a.Address == v.value:
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Email>", *v.tag))
		default:
			return errors.New("failed <String> validation for <Email>")
		}
	})
	return v
}

func (v *ValidatableString) Regex(regex string, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		switch {
		case regexp.MustCompile(regex).MatchString(v.value):
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Regex(%s)>", *v.tag, regex))
		default:
			return errors.New("failed <String> validation for <Regex(" + regex + ")>")
		}
	})
	return v
}

func (v *ValidatableString) Custom(rule func(s string) bool, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		switch {
		case rule(v.value):
			return nil
		case len(msg) > 0:
			return errors.New(msg[0])
		case v.tag != nil:
			return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Custom>", *v.tag))
		default:
			return errors.New("failed <String> validation for <Custom>")
		}
	})
	return v
}

func String() *ValidatableString {
	return &ValidatableString{}
}
