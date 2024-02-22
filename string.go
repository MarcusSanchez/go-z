package z

import (
	"fmt"
	"github.com/MarcusSanchez/go-z/internal"
	"net/mail"
	"regexp"
)

var _ Validatable = (*ValidatableString)(nil)

// ValidatableString is a string that can be validated.
type ValidatableString struct {
	tag      *string
	value    string
	rules    []rule
	optional bool
}

// Validate validates a string against its schema.
//
//	Returns Errors if:
//	=> data is not a string
//	=> data fails any of the schema's rules
func (v *ValidatableString) Validate(data any, tag ...string) Errors {
	if len(tag) > 0 {
		v.tag = &tag[0]
	}
	if v.optional {
		if data == nil {
			return nil
		}
		if value, ok := data.(*string); ok {
			if value == nil {
				return nil
			}
			data = *value
		}
	}
	var ok bool
	if v.value, ok = data.(string); !ok {
		if v.tag == nil {
			return internal.NewValidationErrors("failed validation for <string>")
		}
		return internal.NewValidationErrors(fmt.Sprintf("<%s> failed validation for <string>", *v.tag))
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

// Optional marks the string as optional. Calling Validate with nil or a nil string pointer will skip validation.
func (v *ValidatableString) Optional() *ValidatableString {
	v.optional = true
	return v
}

// Min appends a rule validating that data is greater than or equal to the provided min. (len(data) >= min)
func (v *ValidatableString) Min(min int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		switch {
		case len(v.value) >= min:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <string> validation for <Min(%d)>", *v.tag, min)
		default:
			return fmt.Sprintf("failed <string> validation for <Min(%d)>", min)
		}
	})
	return v
}

// Max appends a rule validating that data is less than or equal to the provided max. (len(data) <= max)
func (v *ValidatableString) Max(max int, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		switch {
		case len(v.value) <= max:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <string> validation for <Max(%d)>", *v.tag, max)
		default:
			return fmt.Sprintf("failed <string> validation for <Max(%d)>", max)
		}
	})
	return v
}

// Email appends a rule validating that data is a valid email address under RFC-5322.
func (v *ValidatableString) Email(msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		a, err := mail.ParseAddress(v.value)
		switch {
		case err == nil && a.Address == v.value:
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <string> validation for <Email>", *v.tag)
		default:
			return "failed <string> validation for <Email>"
		}
	})
	return v
}

// In appends a rule validating that data is in the provided slice of values.
func (v *ValidatableString) In(values []string, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		for _, value := range values {
			if v.value == value {
				return ""
			}
		}
		if len(msg) > 0 {
			return msg[0]
		}
		if v.tag != nil {
			return fmt.Sprintf("<%s> failed <%T> validation for <In(%s)>", *v.tag, v.value, values)
		}
		return fmt.Sprintf("failed <%T> validation for <In(%s)>", v.value, values)
	})
	return v
}

// Regex appends a rule validating that data matches the provided regex.
func (v *ValidatableString) Regex(regex string, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		switch {
		case regexp.MustCompile(regex).MatchString(v.value):
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <string> validation for <Regex(%s)>", *v.tag, regex)
		default:
			return "failed <string> validation for <Regex(" + regex + ")>"
		}
	})
	return v
}

// Custom appends a custom rule to the schema. Validates if the provided function returns true when passed data.
func (v *ValidatableString) Custom(rule func(s string) bool, msg ...string) *ValidatableString {
	v.rules = append(v.rules, func() string {
		switch {
		case rule(v.value):
			return ""
		case len(msg) > 0:
			return msg[0]
		case v.tag != nil:
			return fmt.Sprintf("<%s> failed <string> validation for <Custom>", *v.tag)
		default:
			return "failed <string> validation for <Custom>"
		}
	})
	return v
}

// String returns a new ValidatableString for validation a string.
func String() *ValidatableString { return &ValidatableString{} }
