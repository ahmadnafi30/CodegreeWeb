package entity

type Option struct {
	ID         uint   `json:"id" gorm:"primary_key;autoIncrement"`
	QuestionID uint   `json:"question_id"`
	Value      string `json:"value" gorm:"type:text"`
}
