package schemas

type ThemeSchema interface {
	ThemeName() string
	TemplateDetails() FileSchema
	PDFDetails() PDFSchema
}

type FileSchema struct {
	ThemeDirectory string
	HTMLTemplate   string
	CSSFile        string
}

type PDFSchema struct {
	MarginTop    float64
	MarginBottom float64
	MarginLeft   float64
	MarginRight  float64
	Scale        float64
}
