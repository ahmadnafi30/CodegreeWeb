package entity

type Recom struct {
	ID             uint   `gorm:"primaryKey" json:"-"`
	LanguageCodeID uint   `json:"language_code_id"`
	Name           string `json:"name"`
}
