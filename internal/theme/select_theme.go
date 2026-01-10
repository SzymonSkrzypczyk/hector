package theme

import (
	"hector/internal/schemas"
	"log"
	"strings"
)

func SelectTheme(themeName string) schemas.ThemeSchema {
	switch strings.TrimSpace(strings.ToLower(themeName)) {
	case schemas.Normal{}.ThemeName():
		return &schemas.Normal{}
	case schemas.Gray{}.ThemeName():
		return &schemas.Gray{}
	default:
		log.Println("Unknown theme, defaulting to 'normal'")
		return &schemas.Normal{}
	}
}
