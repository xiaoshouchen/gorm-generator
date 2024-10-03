package model

import (
	"fmt"
	"reflect"
	"strings"
)

func formatInterface(v interface{}) string {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		var elements []string
		for i := 0; i < value.Len(); i++ {
			elements = append(elements, formatInterface(value.Index(i).Interface()))
		}
		return strings.Join(elements, "_")
	case reflect.Map:
		var pairs []string
		for _, key := range value.MapKeys() {
			pair := fmt.Sprintf("%v:%v", formatInterface(key.Interface()), formatInterface(value.MapIndex(key).Interface()))
			pairs = append(pairs, pair)
		}
		return strings.Join(pairs, "_")
	case reflect.Struct:
		var fields []string
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			fieldValue := formatInterface(value.Field(i).Interface())
			fields = append(fields, fmt.Sprintf("%s:%s", field.Name, fieldValue))
		}
		return strings.Join(fields, "_")
	case reflect.Ptr:
		if value.IsNil() {
			return "nil"
		}
		return formatInterface(value.Elem().Interface())
	default:
		return fmt.Sprintf("%v", v)
	}
}
