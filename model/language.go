package model

type Language struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
