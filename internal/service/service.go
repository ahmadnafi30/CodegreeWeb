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
	return &Service{
		UserService:         userService,
		OnBoardingService:   onBoardingService,
		LanguageCodeService: languageService,
	}
}
