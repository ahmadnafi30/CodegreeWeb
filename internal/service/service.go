package service

import (
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/pkg/bcrypt"
	"CodegreeWebbs/pkg/jwt"
	// "CodegreeWebbs/pkg/jwt"
)

type Service struct {
	UserService         IUserService
	OnBoardingService   IOnBoardingService
	LanguageCodeService SLanguage
	PaymentService      SPayment
	CourseService       Scourse
}

type InitParam struct {
	Repository *repository.Repository
	JwtAuth    jwt.Interface
	Bcrypt     bcrypt.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JwtAuth)
	onBoardingService := NewOnBoardingService(param.Repository.OnBoardingRepo)
	languageService := NewLanguageService(param.Repository.LangauageRepo)
	paymentService := NewPaymentService(param.Repository.PaymentRepo)
	courseService := NewCourseService(param.Repository.CourseRepo)

	return &Service{
		UserService:         userService,
		OnBoardingService:   onBoardingService,
		LanguageCodeService: languageService,
		PaymentService:      paymentService,
		CourseService:       courseService,
	}
}
