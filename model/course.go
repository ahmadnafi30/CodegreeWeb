package model

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
	Question string   `json:"question"`
	Options  []string `json:"options"`
}
