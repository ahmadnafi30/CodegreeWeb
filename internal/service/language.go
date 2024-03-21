package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/model"
	"errors"
)

type SLanguage interface {
	CreateLanguage(language entity.LanguageCode) error
	CreateMentor(mentor entity.Mentor) error
	GetAllMentor() ([]model.Mentor, error)
	SelectMentorWhatsAppLink(id uint) (string, error)
}

type LanguageService struct {
	LanguageRepo repository.RLanguage
}

func NewLanguageService(languageRepo repository.RLanguage) SLanguage {
	return &LanguageService{
		LanguageRepo: languageRepo,
	}
}

func (s *LanguageService) CreateLanguage(language entity.LanguageCode) error {
	_, err := s.LanguageRepo.SaveLanguage(language)
	if err != nil {
		return errors.New("failed to create language: " + err.Error())
	}
	return nil
}

func (s *LanguageService) CreateMentor(mentor entity.Mentor) error {
	_, err := s.LanguageRepo.CreateMentor(mentor)
	if err != nil {
		return errors.New("failed to add mentor " + err.Error())
	}
	return nil
}

func (s *LanguageService) GetAllMentor() ([]model.Mentor, error) {
	mentors, err := s.LanguageRepo.GetALLMentor()
	if err != nil {
		return nil, err
	}
	result := make([]model.Mentor, len(mentors))
	for i, v := range mentors {
		result[i] = model.Mentor{
			ID:          v.ID,
			Name:        v.Name,
			Language:    v.Language,
			Description: v.Description,
			Company:     v.Company,
		}
	}
	return result, nil
}

func (s *LanguageService) SelectMentorWhatsAppLink(id uint) (string, error) {
	linkWhatsApp, err := s.LanguageRepo.SelectMentorWhatsAppLink(id)
	if err != nil {
		return "", err
	}

	return linkWhatsApp, nil
}
