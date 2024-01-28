package z

import (
	"fmt"
	"github.com/MarcusSanchez/go-z/internal"
)

var _ Validatable = (*ValidatableBool)(nil)

// ValidatableBool is a bool that can be validated.
type ValidatableBool struct {
	tag      *string
	value    bool
	rules    []rule
	optional bool
}

// Validate validates a bool against its schema.
//
//	Returns Errors if:
//	=> data is not a bool
//	=> data fails any of the schema's rules
func (v *ValidatableBool) Validate(data any, tag ...string) Errors {
	if len(tag) > 0 {
		v.tag = &tag[0]
	}
	if v.optional {
		if data == nil {
			return nil
		}
		if value, ok := data.(*bool); ok {
			if value == nil {
				return nil
			}
			data = *value
		}
	}

	var ok bool
	if v.value, ok = data.(bool); !ok {
		if v.tag == nil {
			return internal.NewValidationErrors("failed validation for <bool>")
		}
		return internal.NewValidationErrors("<%s> failed validation for <bool>", *v.tag)
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

// Optional marks the bool as optional. Calling Validate with nil or a nil bool pointer will skip validation.
func (v *ValidatableBool) Optional() *ValidatableBool {
	v.optional = true
	return v
}

// True appends a rule validating that data is true. (data == true)
func (v *ValidatableBool) True(msg ...string) *ValidatableBool {
	v.rules = append(v.rules, func() string {
		if v.value == true {
			return ""
		}
		if len(msg) > 0 {
			return msg[0]
		}
		if v.tag == nil {
			return fmt.Sprintf("failed <bool> validation for <True>")
		}
		return fmt.Sprintf("<%s> failed <bool> validation for <True>", *v.tag)
	})
	return v
}

// False appends a rule validating that data is false. (data == false)
func (v *ValidatableBool) False(msg ...string) *ValidatableBool {
	v.rules = append(v.rules, func() string {
		if v.value == false {
			return ""
		}
		if len(msg) > 0 {
			return msg[0]
		}
		if v.tag == nil {
			return fmt.Sprintf("failed <bool> validation for <False>")
		}
		return fmt.Sprintf("<%s> failed <bool> validation for <False>", *v.tag)
	})
	return v
}

// Bool returns a ValidatableBool for validating a bool.
func Bool() *ValidatableBool { return &ValidatableBool{} }
