package schemas

type Normal struct {
	CandidateInfo   NormalCandidateInfo `yaml:"candidate_info" json:"candidate_info"`
	Contact         NormalContact       `yaml:"contact" json:"contact"`
	Skills          []string            `yaml:"skills" json:"skills"`
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

type NormalCandidateInfo struct {
	Name  string `yaml:"name" json:"name"`
	About string `yaml:"about" json:"about"`
}

type NormalContact struct {
	Email string `yaml:"email" json:"email"`
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
