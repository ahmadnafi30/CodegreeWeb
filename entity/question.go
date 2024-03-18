package entity

type Question struct {
	ID            uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	SubLanguageID uint     `json:"sub_language_id"`
	Question      string   `json:"question" gorm:"type:text"`
	Options       []Option `json:"options,omitempty"`
}
