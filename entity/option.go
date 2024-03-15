package entity

type Option struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Option     string `json:"option"`
	Value      bool   `json:"value"`
}
