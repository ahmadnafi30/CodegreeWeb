package entity

type Question struct {
	ID            uint     `json:"id"`
	SubLanguageID uint     `json:"sub_language_id"`
	Question      string   `json:"question"`
	Options       []Option `json:"options,omitempty"`
	Score         int64    `json:"score"`
}
