package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"

	// "CodegreeWebbs/model"

	// "CodegreeWebbs/model"
	"errors"

	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type IOnBoardingService interface {
	SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error
	SaveMultipleChoiceOnboarding(onboarding *entity.Onboarding, options []entity.OptionBoarding) error
	GetAllOnboardingQuestions() ([]entity.Onboarding, error)
	GetUserByID(userID uuid.UUID) (entity.User, error)
	CheckAnswer(userId uuid.UUID) ([]entity.LanguageCode, error)
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

func (u *OnBoardingService) GetUserByID(userID uuid.UUID) (entity.User, error) {
	return u.OnBoardingRepo.GetUserByID(userID)
}

func (s *OnBoardingService) SaveUserAnswer(answer *entity.UserAnswerOnBoarding) error {
	if answer.QuestionID == 0 {
		return errors.New("invalid QuestionID")
	}
	if answer.Answer == 0 {
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

func (s *OnBoardingService) CheckAnswer(userId uuid.UUID) ([]entity.LanguageCode, error) {
	var recommendedLanguage entity.LanguageCode
	answer, err := s.OnBoardingRepo.CheckAnswer(userId)
	if err != nil {
		return nil, err
	}

	var languages []entity.LanguageCode

	for _, ans := range answer {
		if ans.QuestionID == 1 {
			switch ans.Answer {
			case 1:
				recommendedLanguage = entity.LanguageCode{ID: 1, Title: "Java", Description: "Bahasa Java adalah bahasa pemrograman tingkat tinggi yang bersifat berbasis objek, yang dikembangkan oleh Sun Microsystems (sekarang dimiliki oleh Oracle Corporation). Bahasa ini didesain agar bersifat mudah dipelajari, ditulis, dibaca, dan dijalankan oleh mesin virtual Java (JVM), yang membuatnya dapat dijalankan pada berbagai platform yang mendukung JVM, seperti Windows, macOS, Linux, dan sebagainya."}
			case 2:
				recommendedLanguage = entity.LanguageCode{ID: 2, Title: "Javascript", Description: "JavaScript adalah bahasa pemrograman tingkat tinggi yang sering digunakan untuk mengembangkan aplikasi web interaktif. Meskipun memiliki nama yang mirip dengan Java, JavaScript sebenarnya adalah bahasa yang berbeda dan memiliki fitur serta sintaksis yang berbeda pula. JavaScript awalnya dikembangkan oleh Netscape Communications Corporation pada tahun 1995 dengan nama LiveScript, kemudian berganti nama menjadi JavaScript."}
			case 3:
				recommendedLanguage = entity.LanguageCode{ID: 3, Title: "Fluter", Description: "Flutter adalah sebuah framework pengembangan aplikasi mobile open-source yang dikembangkan oleh Google. Diperkenalkan pada tahun 2018, Flutter dirancang untuk memungkinkan pengembang untuk membuat aplikasi yang kaya dan indah secara konsisten di berbagai platform seperti Android, iOS, dan bahkan web."}
			}
		}
	}

	if recommendedLanguage.ID != 0 {
		languages = append(languages, recommendedLanguage)
	}

	return languages, nil
}
