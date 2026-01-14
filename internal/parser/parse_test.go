package parser

import (
	"hector/internal/schemas"
	"os"
	"testing"
)

// MockSchema implements schemas.ThemeSchema for testing
type MockSchema struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func (m *MockSchema) ThemeName() string                         { return "mock" }
func (m *MockSchema) TemplateDetails() schemas.FileSchema       { return schemas.FileSchema{} }
func (m *MockSchema) PDFDetails() schemas.PDFSchema             { return schemas.PDFSchema{} }
func (m MockSchema) Validate() (bool, []schemas.ValidationFlaw) { return true, nil }

func TestParse(t *testing.T) {
	content := []byte(`
name: Hector User
description: Software Engineer
`)
	tmpFile, err := os.CreateTemp("", "test_values_*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	schema := &MockSchema{}

	result := Parse(tmpFile.Name(), schema)

	parsed, ok := result.(*MockSchema)
	if !ok {
		t.Fatalf("Expected *MockSchema, got %T", result)
	}

	if parsed.Name != "Hector User" {
		t.Errorf("Expected Name 'Hector User', got '%s'", parsed.Name)
	}
	if parsed.Description != "Software Engineer" {
		t.Errorf("Expected Description 'Software Engineer', got '%s'", parsed.Description)
	}
}
