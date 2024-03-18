package entity

type Option struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	QuestionID uint   `json:"question_id"`
	Option     string `json:"option" gorm:"type:text"`
	Value      bool   `json:"value"`
}
