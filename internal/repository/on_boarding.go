package repository

import (
	"CodegreeWebbs/entity"

	"gorm.io/gorm"
)

type IOnBoarding interface {
	SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error
	SaveOnboarding(onboarding *entity.Onboarding) error
	SaveOption(option *entity.OptionBoarding) error
	GetAllOnboardingQuestions() ([]entity.Onboarding, error)
}

type OnBoardingRepo struct {
	db *gorm.DB
}

func NewOnBoardingRepo(db *gorm.DB) IOnBoarding {
	return &OnBoardingRepo{db: db}
}

func (repo *OnBoardingRepo) SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error {
	if err := repo.db.Create(answer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OnBoardingRepo) SaveOnboarding(onboarding *entity.Onboarding) error {
	if err := repo.db.Create(onboarding).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OnBoardingRepo) SaveOption(option *entity.OptionBoarding) error {
	if err := repo.db.Create(option).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OnBoardingRepo) GetAllOnboardingQuestions() ([]entity.Onboarding, error) {
	var questions []entity.Onboarding
	if err := repo.db.Preload("Options").Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}
