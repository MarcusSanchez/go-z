package z

import (
	"github.com/MarcusSanchez/go-z/internal"
	"reflect"
	"slices"
)

var _ Validatable = (Struct)(nil)

// Struct is a map of z-tags to Validatable schemas.
// Corresponding tags in the struct will be validated against their schemas.
type Struct map[string]Validatable

// Validate validates a struct or struct pointer against its schema.
//
//	Returns Errors if:
//	=> data is not a struct or a (non-nil) struct pointer
//	=> a tag is not found in the schema
//	=> a tag is found in the schema but fails its schema's validation
func (s Struct) Validate(data any, tags ...string) Errors {
	if data == nil {
		if len(tags) > 0 {
			return internal.NewValidationErrors("<" + tags[0] + "> failed validation for <struct> (nil interface)")
		}
		return internal.NewValidationErrors("failed validation for <struct> (nil interface)")
	}

	// ensure data is a struct or struct pointer
	t := reflect.TypeOf(data)
	kind := t.Kind()
	value := reflect.ValueOf(data)

	if kind == reflect.Ptr {
		if value.IsNil() {
			// if data is a nil pointer, and struct isn't optional, return an error
			if len(tags) > 0 {
				return internal.NewValidationErrors("<" + tags[0] + "> failed validation for <struct> (nil pointer)")
			}
			return internal.NewValidationErrors("failed validation for <struct> (nil pointer)")
		}
		// if data is a pointer, dereference it
		data = value.Elem().Interface()
		t = reflect.TypeOf(data)
		kind = t.Kind()
	}
	if kind != reflect.Struct {
		// if data is not a struct, even after dereferencing, return an error
		if len(tags) > 0 {
			return internal.NewValidationErrors("failed validation for <" + tags[0] + ">")
		}
		return internal.NewValidationErrors("failed validation for <struct>")
	}

	// grab the values from the struct
	values := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("z")

		if tag == "" || tag == "-" {
			continue
		}
		values[tag] = reflect.ValueOf(data).FieldByName(field.Name).Interface()
	}

	// due to maps being unordered, sort tags to allow for predictable validation
	keys := make([]string, 0, len(s))
	for k := range values {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	errs := internal.NewValidationErrors()
	for _, tag := range keys {
		schema := s[tag]

		value, exists := values[tag]
		if len(tags) > 0 {
			tag = tags[0] + "." + tag
		}
		if !exists {
			// if there's no matching value, don't bother validating
			return internal.NewValidationErrors("tag <" + tag + "> not found for <struct>")
		}

		// recursively validate values, appending any errors to the returned ValidationErrors
		if err := schema.Validate(value, tag); err != nil {
			errs.Errors = append(errs.Errors, err.All()...)
		}
	}

	if len(errs.Errors) > 0 {
		return errs
	}
	return nil
}

func (s Struct) Optional() OptionalStruct {
	return OptionalStruct(s)
}

type OptionalStruct map[string]Validatable

// Validate validates a struct or a struct pointer (if it's not nil) against its schema.
//
//	Returns Errors if:
//	=> data is not a struct or a struct pointer
//	=> a tag is not found in the schema
//	=> a tag is found in the schema but fails its schema's validation
func (s OptionalStruct) Validate(data any, tags ...string) Errors {
	if data == nil {
		return nil
	}

	t := reflect.TypeOf(data)
	kind := t.Kind()
	value := reflect.ValueOf(data)

	if kind == reflect.Ptr {
		if value.IsNil() {
			return nil
		}
	}

	return Struct(s).Validate(data, tags...)
}
