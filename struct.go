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
	default:
		switch len(tags) {
		case 0:
			return errors.New("failed validation for <Struct>")
		default:
			return errors.New(tags[0] + " failed validation for <Struct>")
		}
	}
	t := reflect.TypeOf(data)

	values := make(map[string]any, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		values[t.Field(i).Tag.Get("z")] = reflect.ValueOf(data).FieldByName(t.Field(i).Name).Interface()
	}

	var errs []string
	for zTag, value := range s {
		v, ok := values[zTag]
		if !ok {
			return errors.New("tag <" + zTag + "> not found for <Struct>")
		}

		if len(tags) > 0 {
			zTag = tags[0] + "." + zTag
		}
		if err := value.Validate(v, zTag); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}
