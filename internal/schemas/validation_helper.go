package schemas

import (
	"fmt"
	"reflect"
)

func isEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Slice, reflect.Array, reflect.Map:
		return val.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	case reflect.Struct:
		return false
	default:
		return false
	}
}

func ValidateRequiredFields(s interface{}) []string {
	var missing []string
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return missing
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "required" && isEmpty(field) {
			missing = append(missing, fieldType.Name)
		}

		if field.Kind() == reflect.Struct {
			nested := ValidateRequiredFields(field.Interface())
			for _, n := range nested {
				missing = append(missing, fieldType.Name+"."+n)
			}
		}

		if field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Struct {
			for j := 0; j < field.Len(); j++ {
				nested := ValidateRequiredFields(field.Index(j).Interface())
				for _, n := range nested {
					missing = append(missing, fieldType.Name+fmt.Sprintf("[%d].%s", j, n))
				}
			}
		}
	}

	return missing
}
