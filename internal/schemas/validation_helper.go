package schemas

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	phoneField = "phone"
	emailField = "email"
)

type ValidationFlaw struct {
	kind string
	name string
}

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

func ValidatePhone(phoneNumber string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(phoneNumber)
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func GroupByKind(flaws []ValidationFlaw) map[string][]string {
	grouped := make(map[string][]string)
	for _, f := range flaws {
		grouped[f.kind] = append(grouped[f.kind], f.name)
	}
	return grouped
}

func ValidateFields(s interface{}) []ValidationFlaw {
	var validation_flaws []ValidationFlaw
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return validation_flaws
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return validation_flaws
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		validate_tag := fieldType.Tag.Get("validate")
		max_len_tag := fieldType.Tag.Get("max_len")

		if validate_tag == "required" && isEmpty(field) {
			validation_flaws = append(validation_flaws,
				ValidationFlaw{
					"required",
					fieldType.Name,
				})
		}

		if max_len_tag != "" {
			converted_max_len, err := strconv.Atoi(max_len_tag)
			if err != nil {
				log.Fatal(err)
			}

			if field.Kind() == reflect.String && field.Len() > converted_max_len {
				validation_flaws = append(validation_flaws, ValidationFlaw{
					"max_len",
					fieldType.Name,
				})
			}
		}

		if strings.Contains(strings.ToLower(fieldType.Name), phoneField) {
			if !ValidatePhone(field.String()) {
				validation_flaws = append(validation_flaws, ValidationFlaw{
					"wrong_format",
					fieldType.Name,
				})
			}
		}

		if strings.Contains(strings.ToLower(fieldType.Name), emailField) {
			if !ValidateEmail(field.String()) {
				validation_flaws = append(validation_flaws, ValidationFlaw{
					"wrong_format",
					fieldType.Name,
				})
			}
		}

		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && !field.IsNil()) {
			nested := ValidateFields(field.Interface())
			for _, n := range nested {
				validation_flaws = append(validation_flaws,
					ValidationFlaw{
						n.kind,
						fieldType.Name + "." + n.name,
					})
			}
		}

		isSlice := field.Kind() == reflect.Slice
		if isSlice {
			elemKind := field.Type().Elem().Kind()
			// Allow recursion if slice contains Structs OR Pointers
			if elemKind == reflect.Struct || elemKind == reflect.Ptr {
				for j := 0; j < field.Len(); j++ {
					nested := ValidateFields(field.Index(j).Interface())
					for _, n := range nested {
						validation_flaws = append(validation_flaws,
							ValidationFlaw{
								n.kind,
								fieldType.Name + fmt.Sprintf("[%d].%s", j, n),
							})
					}
				}
			}
		}
	}

	return validation_flaws
}
