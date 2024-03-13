package repository

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	SeeProfile(userID string) (model.UserProfile, error)
	CreateProfile(profile model.UserProfile) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) CreateProfile(profile model.UserProfile) error {
	err := u.db.Debug().Create(&profile).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) SeeProfile(userID string) (model.UserProfile, error) {
	var userProfile model.UserProfile
	err := u.db.Debug().Where("id = ?", userID).First(&userProfile).Error
	if err != nil {
		return model.UserProfile{}, err
	}
	return userProfile, nil
}
