package entity

type Course struct {
	ID           uint          `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	LanguageID   uint          `json:"language_id"`
	Language     LanguageCode  `json:"language,omitempty"`
	SubLanguages []SubLanguage `json:"sub_languages,omitempty"`
}
