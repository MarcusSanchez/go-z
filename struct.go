package z

import (
	"errors"
	"reflect"
	"strings"
)

var _ Validatable = (Struct)(nil)

type Struct map[string]Validatable

func (s Struct) Validate(data any, tags ...string) error {

	switch reflect.TypeOf(data).Kind() {
	case reflect.Struct:
		break // do nothing
	case reflect.Ptr:
		data = reflect.ValueOf(data).Elem().Interface()
		if reflect.TypeOf(data).Kind() == reflect.Struct {
			break
		}
		fallthrough
	default:
		if len(tags) > 0 {
			return errors.New(tags[0] + " failed validation for <Struct>")
		}
		return errors.New("failed validation for <Struct>")
	}
	t := reflect.TypeOf(data)

	values := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		values[t.Field(i).Tag.Get("z")] = reflect.ValueOf(data).FieldByName(t.Field(i).Name).Interface()
	}

	var errs []string
	for tag, schema := range s {
		value, exists := values[tag]
		if len(tags) > 0 {
			tag = tags[0] + "." + tag
		}
		if !exists {
			return errors.New("tag <" + tag + "> not found for <Struct>")
		}
		if err := schema.Validate(value, tag); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}
