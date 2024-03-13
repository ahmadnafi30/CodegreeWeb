package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"errors"
	// "github.com/google/uuid"
)

type IOnBoardingService interface {
	SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error
	SaveMultipleChoiceOnboarding(onboarding *entity.Onboarding, options []entity.OptionBoarding) error
	GetAllOnboardingQuestions() ([]entity.Onboarding, error)
}

type OnBoardingService struct {
	OnBoardingRepo repository.IOnBoarding
}

func NewOnBoardingService(onBoardingRepo repository.IOnBoarding) IOnBoardingService {
	return &OnBoardingService{OnBoardingRepo: onBoardingRepo}
}

func (s *OnBoardingService) SaveMultipleChoiceOnboarding(onboarding *entity.Onboarding, options []entity.OptionBoarding) error {
	if err := s.OnBoardingRepo.SaveOnboarding(onboarding); err != nil {
		return err
	}
	for i := range options {
		options[i].OnboardingID = onboarding.ID
		if err := s.OnBoardingRepo.SaveOption(&options[i]); err != nil {
			return err
		}
	}

	return nil
}

func (s *OnBoardingService) SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error {
	if answer.QuestionID == 0 {
		return errors.New("invalid QuestionID")
	}
	if answer.Answer == "" {
		return errors.New("empty Response")
	}
	return s.OnBoardingRepo.SaveUserAnswer(answer)
}

func (s *OnBoardingService) GetAllOnboardingQuestions() ([]entity.Onboarding, error) {
	questions, err := s.OnBoardingRepo.GetAllOnboardingQuestions()
	if err != nil {
		return nil, err
	}

	return questions, nil
}
