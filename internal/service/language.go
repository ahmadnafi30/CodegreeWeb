package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"errors"
)

type SLanguage interface {
	CreateLanguage(language entity.LanguageCode) error
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
