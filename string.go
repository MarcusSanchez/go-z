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

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)

func (v *ValidatableString) Email(msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		if emailRegex.MatchString(v.value) {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		if v.tag == nil {
			return errors.New("failed <String> validation for <Email>")
		}
		return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Email>", *v.tag))
	})
	return v
}

func (v *ValidatableString) Min(min int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		if !(len(v.value) < min) {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed <String> validation for <Min(%d)>", min))
		}
		return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Min(%d)>", *v.tag, min))
	})
	return v
}

func (v *ValidatableString) Max(max int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() error {
		if !(len(v.value) > max) {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed <String> validation for <Max(%d)>", max))
		}
		return errors.New(fmt.Sprintf("<%s> failed <String> validation for <Max(%d)>", *v.tag, max))
	})
	return v
}

func String() *ValidatableString {
	return &ValidatableString{}
}
