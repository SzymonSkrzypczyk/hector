package render

import (
	"hector/internal/schemas"
)

func ExtractFiles(theme schemas.ThemeSchema) schemas.FileSchema {
	switch t := theme.(type) {
	case *schemas.Normal:
		return t.TemplateDetails()
	default:
		return schemas.FileSchema{}
	}
}
