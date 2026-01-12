package theme

import (
	"testing"
)

func TestSelectTheme(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
	}{
		{
			name:     "Select Normal",
			input:    "normal",
			wantName: "normal",
		},
		{
			name:     "Select Gray",
			input:    "gray",
			wantName: "gray",
		},
		{
			name:     "Case Insensitive",
			input:    "GRAY",
			wantName: "gray",
		},
		{
			name:     "Unknown defaults to Normal",
			input:    "super-theme",
			wantName: "normal",
		},
		{
			name:     "Empty defaults to Normal",
			input:    "",
			wantName: "normal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectTheme(tt.input)
			if got.ThemeName() != tt.wantName {
				t.Errorf("SelectTheme(%q) = %v, want %v", tt.input, got.ThemeName(), tt.wantName)
			}
		})
	}
}
