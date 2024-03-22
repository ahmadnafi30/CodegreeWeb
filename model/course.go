package model

// import "CodegreeWebbs/model"

type GetCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetCoursedetail struct {
	Progress int64    `json:"progress"`
	Sublang  []string `json:"sublang"`
}

type GetSublang struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Material    string `json:"material"`
}

type Gamification struct {
	QuestionID uint
	Question   string `json:"question"`
	Options    []Option
}

type Option struct {
	ID     uint
	Option string
}

type Mentor struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	Name        string `json:"name"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Company     string `json:"company"`
}
