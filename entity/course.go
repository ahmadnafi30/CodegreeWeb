package entity

type Course struct {
	ID           uint          `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	SubLanguages []SubLanguage `json:"sub_languages,omitempty"`
}
