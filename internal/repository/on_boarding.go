package repository

import (
	"CodegreeWebbs/entity"
	// "CodegreeWebbs/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IOnBoarding interface {
	SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error
	SaveOnboarding(onboarding *entity.Onboarding) error
	SaveOption(option *entity.OptionBoarding) error
	GetAllOnboardingQuestions() ([]entity.Onboarding, error)
	GetUserByID(userID uuid.UUID) (entity.User, error)
	CheckAnswer(userId uuid.UUID) ([]entity.UserAnswerOnBoarding, error)
	// GetUnansweredQuestions(userID int) ([]entity.Onboarding, error)
}

type OnBoardingRepo struct {
	db *gorm.DB
}

func NewOnBoardingRepo(db *gorm.DB) IOnBoarding {
	return &OnBoardingRepo{db: db}
}

func (u *OnBoardingRepo) GetUserByID(userID uuid.UUID) (entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("id = ?", userID).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
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

func (repo *OnBoardingRepo) CheckAnswer(userID uuid.UUID) ([]entity.UserAnswerOnBoarding, error) {
	var check []entity.UserAnswerOnBoarding

	if err := repo.db.Where("user_id = ? AND question_id = ?", userID, 1).Find(&check).Error; err != nil {
		return nil, err
	}
	return check, nil
}

// func (repo *OnBoardingRepo) CheckAnswer()([]entity.UserAnswerOnBoarding, error) {
// 	var answer []entity.UserAnswerOnBoarding

// }
