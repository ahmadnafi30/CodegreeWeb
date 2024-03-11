package model

// import "github.com/google/uuid"

type CreateMultipleChoiceOnboarding struct {
	Question string `json:"question"`
	Options  []struct {
		Description string `json:"description"`
		Language    string `json:"-"`
	} `json:"options"`
}
