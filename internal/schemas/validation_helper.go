package schemas

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func parsePeriod(period string) (start, end *time.Time, err error) {
	parts := strings.Split(period, "-")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid period format: %s", period)
	}

	startStr := strings.TrimSpace(parts[0])
	endStr := strings.TrimSpace(parts[1])

	layout := "January 2006" // Example: "March 2024"
	startTime, err := time.Parse(layout, startStr)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse start date: %v", err)
	}

	var endTime *time.Time
	if strings.EqualFold(endStr, "Onwards") {
		endTime = nil // nil means ongoing
	} else {
		et, err := time.Parse(layout, endStr)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot parse end date: %v", err)
		}
		endTime = &et
	}

	return &startTime, endTime, nil
}

func validateChronologicalPeriods(periods []string) []ValidationFlaw {
	var flaws []ValidationFlaw
	var prevStart *time.Time

	for i, p := range periods {
		start, end, err := parsePeriod(p)
		if err != nil {
			flaws = append(flaws,
				ValidationFlaw{
					"invalid_period",
					fmt.Sprintf("period[%d]", i),
				})
			continue
		}

		if end != nil && start.After(*end) {
			flaws = append(flaws,
				ValidationFlaw{
					"invalid_period",
					fmt.Sprintf("period[%d] (start after end)", i),
				})
		}

		if prevStart != nil && start.Before(*prevStart) {
			flaws = append(flaws,
				ValidationFlaw{
					"not_chronological",
					fmt.Sprintf("period[%d] (earlier than previous element)", i),
				})
		}

		prevStart = start
	}

	return flaws
}

func checkDuplicates(slice reflect.Value) []string {
	seen := make(map[string]bool)
	var duplicates []string
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		if item.Kind() != reflect.String {
			continue
		}
		str := item.String()
		if seen[str] {
			duplicates = append(duplicates, str)
		} else {
			seen[str] = true
		}
	}
	return duplicates
}

func GroupByKind(flaws []ValidationFlaw) map[string][]string {
	grouped := make(map[string][]string)
	for _, f := range flaws {
		grouped[f.kind] = append(grouped[f.kind], f.name)
	}
	return grouped
}

func ValidateFields(s interface{}) []ValidationFlaw {
	var validationFlaws []ValidationFlaw
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return validationFlaws
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return validationFlaws
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		validate_tag := fieldType.Tag.Get("validate")
		max_len_tag := fieldType.Tag.Get("max_len")
		regexTag := fieldType.Tag.Get("expected_regex")

		// Required field check
		if validate_tag == "required" && isEmpty(field) {
			validationFlaws = append(validationFlaws,
				ValidationFlaw{
					"required",
					fieldType.Name,
				})
		}

		// Max length check
		if max_len_tag != "" {
			converted_max_len, err := strconv.Atoi(max_len_tag)
			if err != nil {
				log.Fatal(err)
			}

			if field.Kind() == reflect.String && field.Len() > converted_max_len {
				validationFlaws = append(validationFlaws, ValidationFlaw{
					"max_len",
					fieldType.Name,
				})
			}
		}

		// format check
		if regexTag != "" && field.Kind() == reflect.String {
			re := regexp.MustCompile(regexTag)
			if !re.MatchString(field.String()) {
				validationFlaws = append(validationFlaws, ValidationFlaw{
					"wrong_format",
					fieldType.Name,
				})
			}
		}

		// Nested structs or pointers
		if field.Kind() == reflect.Struct || (field.Kind() == reflect.Ptr && !field.IsNil()) {
			nested := ValidateFields(field.Interface())
			for _, n := range nested {
				validationFlaws = append(validationFlaws,
					ValidationFlaw{
						n.kind,
						fieldType.Name + "." + n.name,
					})
			}
		}

		// Slice handling
		if field.Kind() == reflect.Slice {
			elemKind := field.Type().Elem().Kind()
			if elemKind == reflect.Struct || elemKind == reflect.Ptr {
				var prevStart *time.Time
				for j := 0; j < field.Len(); j++ {
					elemVal := field.Index(j)
					if elemVal.Kind() == reflect.Ptr && !elemVal.IsNil() {
						elemVal = elemVal.Elem()
					}

					// validate nested fields
					nested := ValidateFields(elemVal.Interface())
					for _, n := range nested {
						validationFlaws = append(validationFlaws,
							ValidationFlaw{
								n.kind,
								fieldType.Name + fmt.Sprintf("[%d].%s", j, n.name),
							})
					}

					// Check for "date" fields in this element
					if elemVal.Kind() == reflect.Struct {
						for k := 0; k < elemVal.NumField(); k++ {
							subFieldType := elemVal.Type().Field(k)
							subField := elemVal.Field(k)
							if subFieldType.Tag.Get("validate") == "date" && subField.Kind() == reflect.String {
								start, _, err := parsePeriod(subField.String())
								if err != nil {
									validationFlaws = append(validationFlaws,
										ValidationFlaw{
											"invalid_period",
											fieldType.Name + fmt.Sprintf("[%d].%s", j, subFieldType.Name),
										})
									continue
								}
								if prevStart != nil && start.Before(*prevStart) {
									validationFlaws = append(validationFlaws,
										ValidationFlaw{
											"not_chronological",
											fieldType.Name + fmt.Sprintf("[%d].%s", j, subFieldType.Name),
										})
								}
								prevStart = start
							}
						}
					}
				}
			}

			if elemKind == reflect.String {
				dups := checkDuplicates(field)
				for _, d := range dups {
					validationFlaws = append(validationFlaws,
						ValidationFlaw{
							"duplicate_entry",
							fieldType.Name + " (duplicate: " + d + ")",
						})
				}
			}
		}

		// Also check if the current field itself is tagged with "date" and is a string
		if validate_tag == "date" && field.Kind() == reflect.String {
			dateFlaws := validateChronologicalPeriods([]string{field.String()})
			for _, f := range dateFlaws {
				validationFlaws = append(validationFlaws,
					ValidationFlaw{
						f.kind,
						fieldType.Name,
					})
			}
		}
	}

	return validationFlaws
}
