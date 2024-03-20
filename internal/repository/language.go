package repository

import (
	"CodegreeWebbs/entity"
	"errors"

	"gorm.io/gorm"
)

type RLanguage interface {
	SaveLanguage(language entity.LanguageCode) (entity.LanguageCode, error)
	CreateMentor(mentor entity.Mentor) (entity.Mentor, error)
	GetALLMentor() ([]entity.Mentor, error)
}

type LanguageRepo struct {
	db *gorm.DB
}

func NewLanguageRepo(db *gorm.DB) RLanguage {
	return &LanguageRepo{db: db}
}

func (repo *LanguageRepo) SaveLanguage(language entity.LanguageCode) (entity.LanguageCode, error) {
	err := repo.db.Debug().Create(&language).Error
	if err != nil {
		return entity.LanguageCode{}, errors.New("failed to save language: " + err.Error())
	}
	return language, nil
}

func (repo *LanguageRepo) CreateMentor(mentor entity.Mentor) (entity.Mentor, error) {
	err := repo.db.Debug().Create(&mentor).Error
	if err != nil {
		return entity.Mentor{}, err
	}
	return mentor, nil
}

func (repo *LanguageRepo) GetALLMentor() ([]entity.Mentor, error) {
	var mentors []entity.Mentor
	if err := repo.db.Debug().Find(&mentors).Error; err != nil {
		return nil, err
	}
	return mentors, nil
}
