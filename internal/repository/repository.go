package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	OnBoardingRepo IOnBoarding
	LangauageRepo  RLanguage
	PaymentRepo    IPayment
	CourseRepo     ICourse
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	OnBoardingRepo := NewOnBoardingRepo(db)
	return &Repository{
		UserRepository: userRepository,
		OnBoardingRepo: OnBoardingRepo,
		LangauageRepo:  NewLanguageRepo(db),
		PaymentRepo:    NewPaymentRepo(db),
		CourseRepo:     NewCourseRepo(db),
	}
}
