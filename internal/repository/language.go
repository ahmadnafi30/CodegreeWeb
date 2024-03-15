package repository

import (
	"CodegreeWebbs/entity"
	"errors"

	"gorm.io/gorm"
)

type RLanguage interface {
	SaveLanguage(language entity.LanguageCode) (entity.LanguageCode, error)
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
