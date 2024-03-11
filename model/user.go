package model

import "github.com/google/uuid"

type UserRegister struct {
	ID       uint   `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginAcc struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
}

type UserProfile struct {
	ID     uuid.UUID
	Name   string `json:"name"`
	Email  string `json:"email"`
	Level  int    `json:"level"`
	Xp     uint64 `json:"xp"`
	Hearth int    `json:"hearth"`
}

type UserAnswerOnBoarding struct {
	ID         uint      `json:"id"`
	UserID     uuid.UUID `json:"-"`
	QuestionID uint      `json:"question_id"`
	Response   string    `json:"response"`
}
