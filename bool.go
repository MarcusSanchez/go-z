package z

import (
	"errors"
	"fmt"
	"strings"
)

var _ Validatable = (*ValidatableBool)(nil)

type ValidatableBool struct {
	tag    *string
	value  bool
	rules  []Rule
	errors []string
}

func (v *ValidatableBool) Validate(data any, tag ...string) error {
	if len(tag) > 0 {
		v.tag = &tag[0]
	}
	var ok bool
	if v.value, ok = data.(bool); !ok {
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed validation for <Bool>"))
		}
		return errors.New(fmt.Sprintf("<%s> failed validation for <Bool>", *v.tag))
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

func (v *ValidatableBool) True(msg ...string) *ValidatableBool {
	v.rules = append(v.rules, func() error {
		if v.value == true {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed <Bool> validation for <True>"))
		}
		return errors.New(fmt.Sprintf("<%s> failed <Bool> validation for <True>", *v.tag))
	})
	return v
}

func (v *ValidatableBool) False(msg ...string) *ValidatableBool {
	v.rules = append(v.rules, func() error {
		if v.value == false {
			return nil
		}
		if len(msg) > 0 {
			return errors.New(msg[0])
		}
		if v.tag == nil {
			return errors.New(fmt.Sprintf("failed <Bool> validation for <False>"))
		}
		return errors.New(fmt.Sprintf("<%s> failed <Bool> validation for <False>", *v.tag))
	})
	return v
}

func Bool() *ValidatableBool {
	return &ValidatableBool{}
}
