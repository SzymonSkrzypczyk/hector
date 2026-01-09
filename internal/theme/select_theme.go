package theme

import (
	"hector/internal/schemas"
	"strings"
)

func SelectTheme(themeName string) schemas.ThemeSchema {
	switch strings.TrimSpace(strings.ToLower(themeName)) {
	case schemas.Normal{}.ThemeName():
		return &schemas.Normal{}
	default:
		return &schemas.Normal{}
	}
}
