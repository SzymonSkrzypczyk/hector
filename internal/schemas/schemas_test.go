package schemas

import (
	"reflect"
	"testing"
)

// Define structs for testing purposes
type MockStruct struct {
	Name    string `validate:"required"`
	Age     int
	Address *MockAddress
	Tags    []string `validate:"required"`
}

type MockAddress struct {
	City   string `validate:"required"`
	Street string
}

type NestedSliceStruct struct {
	Items []MockItem
}

type MockItem struct {
	ID string `validate:"required"`
}

func TestValidateRequiredFields(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			name: "Valid struct",
			input: MockStruct{
				Name: "John",
				Tags: []string{"go"},
				Address: &MockAddress{
					City: "NY",
				},
			},
			expected: nil,
		},
		{
			name: "Missing root field",
			input: MockStruct{
				Name: "", // Required
				Tags: []string{"go"},
			},
			expected: []string{"Name"},
		},
		{
			name: "Missing slice field",
			input: MockStruct{
				Name: "John",
				Tags: []string{}, // Empty slice, required
			},
			expected: []string{"Tags"},
		},
		{
			name: "Missing nested pointer field",
			input: MockStruct{
				Name: "John",
				Tags: []string{"go"},
				Address: &MockAddress{
					City: "", // Required
				},
			},
			expected: []string{"Address.City"},
		},
		{
			name: "Missing field in slice of structs",
			input: NestedSliceStruct{
				Items: []MockItem{
					{ID: "1"},
					{ID: ""}, // Required
				},
			},
			expected: []string{"Items[1].ID"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateFields(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				if len(got) == 0 && len(tt.expected) == 0 {
					return
				}
				t.Errorf("ValidateFields() = %v, want %v", got, tt.expected)
			}
		})
	}
}
