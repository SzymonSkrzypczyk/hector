package schemas

type Normal struct {
	CandidateInfo   NormalCandidateInfo `yaml:"candidate_info" json:"candidate_info" validate:"required"`
	Contact         NormalContact       `yaml:"contact" json:"contact" validate:"required"`
	Skills          []string            `yaml:"skills" json:"skills"`
	Photo           string              `yaml:"photo,omitempty" json:"photo,omitempty"`
	Languages       []NormalLanguage    `yaml:"languages" json:"languages"`
	Education       []NormalEducation   `yaml:"education" json:"education"`
	Experience      []NormalExperience  `yaml:"experience" json:"experience"`
	Certifications  []string            `yaml:"certifications" json:"certifications"`
	ExtraActivities []string            `yaml:"extra_activities" json:"extra_activities"`
}

func (Normal) ThemeName() string {
	return "normal"
}

func (Normal) TemplateDetails() FileSchema {
	return FileSchema{
		ThemeDirectory: "normal",
		HTMLTemplate:   "template.html",
		CSSFile:        "style.css",
	}
}

func (Normal) PDFDetails() PDFSchema {
	return PDFSchema{
		MarginTop:    0.0,
		MarginBottom: 0.0,
		MarginLeft:   0.0,
		MarginRight:  0.0,
		Scale:        0.9,
	}
}

func (n *Normal) Validate() (bool, []string) {
	missingFields := ValidateRequiredFields(n)

	if len(missingFields) > 0 {
		return false, missingFields
	}

	return true, nil
}

type NormalCandidateInfo struct {
	Name  string `yaml:"name" json:"name" validate:"required"`
	About string `yaml:"about" json:"about" validate:"required"`
}

type NormalContact struct {
	Email string `yaml:"email" json:"email" validate:"required"`
	Phone string `yaml:"phone" json:"phone"`
}

type NormalLanguage struct {
	Language string `yaml:"language" json:"language"`
	Level    string `yaml:"level" json:"level"`
}

type NormalEducation struct {
	Institution string   `yaml:"institution" json:"institution"`
	Degree      string   `yaml:"degree" json:"degree"`
	Dates       string   `yaml:"period" json:"period"`
	Details     []string `yaml:"details" json:"details"`
}

type NormalExperience struct {
	Role             string   `yaml:"role" json:"role"`
	Company          string   `yaml:"company" json:"company"`
	Dates            string   `yaml:"period" json:"period"`
	Responsibilities []string `yaml:"responsibilities" json:"responsibilities"`
}
