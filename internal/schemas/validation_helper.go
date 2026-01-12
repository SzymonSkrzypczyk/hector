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

	// Unwrap pointer if the input itself is a pointer
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return missing
		}
		val = val.Elem()
	}

	// If it's not a struct (or a pointer to one), we can't validate fields
	if val.Kind() != reflect.Struct {
		return missing
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		// 1. Check current field
		if tag == "required" && isEmpty(field) {
			missing = append(missing, fieldType.Name)
		}

		// 2. Recurse into Nested Structs AND Pointers to Structs
		// FIX: Added check for (reflect.Ptr && !IsNil)
		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && !field.IsNil()) {
			nested := ValidateRequiredFields(field.Interface())
			for _, n := range nested {
				missing = append(missing, fieldType.Name+"."+n)
			}
		}

		// 3. Recurse into Slices
		// We check if the slice holds structs or pointers to structs
		isSlice := field.Kind() == reflect.Slice
		if isSlice {
			elemKind := field.Type().Elem().Kind()
			// Allow recursion if slice contains Structs OR Pointers
			if elemKind == reflect.Struct || elemKind == reflect.Ptr {
				for j := 0; j < field.Len(); j++ {
					nested := ValidateRequiredFields(field.Index(j).Interface())
					for _, n := range nested {
						missing = append(missing, fieldType.Name+fmt.Sprintf("[%d].%s", j, n))
					}
				}
			}
		}
	}

	return missing
}
