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

func (n *Normal) Validate() (bool, []ValidationFlaw) {
	missingFields := ValidateFields(n)

	if len(missingFields) > 0 {
		return false, missingFields
	}

	return true, nil
}

type NormalCandidateInfo struct {
	Name  string `yaml:"name" json:"name" validate:"required" max_len:"100"`
	About string `yaml:"about" json:"about" validate:"required" max_len:"255"`
}

type NormalContact struct {
	Email string `yaml:"email" json:"email" validate:"required" max_len:"100" expected_regex:"^[a-zA-Z0-9._%+\\-]+@[a-zA-Z0-9.\\-]+\\.[a-zA-Z]{2,}$"`
	Phone string `yaml:"phone" json:"phone" max_len:"20" expected_regex:"^\\+?[1-9]\\d{1,14}$"`
}

type NormalLanguage struct {
	Language string `yaml:"language" json:"language"`
	Level    string `yaml:"level" json:"level"`
}

type NormalEducation struct {
	Institution string   `yaml:"institution" json:"institution" max_len:"70"`
	Degree      string   `yaml:"degree" json:"degree"`
	Dates       string   `yaml:"period" json:"period" validate:"date"`
	Details     []string `yaml:"details" json:"details"`
}

type NormalExperience struct {
	Role             string   `yaml:"role" json:"role" max_len:"40"`
	Company          string   `yaml:"company" json:"company"`
	Dates            string   `yaml:"period" json:"period" validate:"date"`
	Responsibilities []string `yaml:"responsibilities" json:"responsibilities"`
}
