package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	OnBoardingRepo IOnBoarding
	LangauageRepo  RLanguage
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)

	return &Repository{
		UserRepository: userRepository,
		OnBoardingRepo: NewOnBoardingRepo(db),
		LangauageRepo:  NewLanguageRepo(db),
	}
}
