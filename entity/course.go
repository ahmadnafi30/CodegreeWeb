package entity

type Course struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Title        string        `json:"title" gorm:"type:varchar(255);not null"`
	Description  string        `json:"description" gorm:"type:text;not null"`
	SubLanguages []SubLanguage `json:"sub_languages,omitempty" gorm:"many2many:course_sub_languages"`
	Progres      int64         `json:"Pogress"`
}
