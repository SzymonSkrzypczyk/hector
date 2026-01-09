package schemas

type Gray struct {
	CandidateInfo   GrayCandidateInfo `yaml:"candidate_info" json:"candidate_info"`
	Contact         GrayContact       `yaml:"contact" json:"contact"`
	Skills          []string          `yaml:"skills" json:"skills"`
	Languages       []GrayLanguage    `yaml:"languages" json:"languages"`
	Education       []GrayEducation   `yaml:"education" json:"education"`
	Experience      []GrayExperience  `yaml:"experience" json:"experience"`
	Certifications  []string          `yaml:"certifications" json:"certifications"`
	ExtraActivities []string          `yaml:"extra_activities" json:"extra_activities"`
}

func (Gray) ThemeName() string {
	return "gray"
}

func (Gray) TemplateDetails() FileSchema {
	return FileSchema{
		ThemeDirectory: "gray",
		HTMLTemplate:   "template.html",
		CSSFile:        "style.css",
	}
}

func (Gray) PDFDetails() PDFSchema {
	return PDFSchema{
		MarginTop:    0.0,
		MarginBottom: 0.0,
		MarginLeft:   0.0,
		MarginRight:  0.0,
		Scale:        0.95,
	}
}

type GrayCandidateInfo struct {
	Name  string `yaml:"name" json:"name"`
	Title string `yaml:"title" json:"title"` // Added Title field
	About string `yaml:"about" json:"about"`
}

type GrayContact struct {
	Email   string `yaml:"email" json:"email"`
	Phone   string `yaml:"phone" json:"phone"`
	Address string `yaml:"address" json:"address"` // Added Address field
}

type GrayLanguage struct {
	Language string `yaml:"language" json:"language"`
	Level    string `yaml:"level" json:"level"`
}

type GrayEducation struct {
	Institution string   `yaml:"institution" json:"institution"`
	Degree      string   `yaml:"degree" json:"degree"`
	Dates       string   `yaml:"dates" json:"dates"` // Mapped to 'dates' from YAML
	Details     []string `yaml:"details" json:"details"`
}

type GrayExperience struct {
	Role             string   `yaml:"role" json:"role"`
	Company          string   `yaml:"company" json:"company"`
	Dates            string   `yaml:"dates" json:"dates"` // Mapped to 'dates' from YAML
	Responsibilities []string `yaml:"responsibilities" json:"responsibilities"`
}
