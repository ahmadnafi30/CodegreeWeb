package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/bcrypt"
	"CodegreeWebbs/pkg/jwt"
	"errors"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	Login(param model.LoginAcc) (model.UserLoginResponse, error)
	GetProfile(userID string) (model.UserProfile, error)
	GetUser(param model.UserParam) (entity.User, error)
}

type UserService struct {
	user   repository.IUserRepository
	bcrypt bcrypt.Interface
	jwt    jwt.Interface
}

func NewUserService(user repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface) IUserService {
	return &UserService{
		user:   user,
		bcrypt: bcrypt,
		jwt:    jwt,
	}
}

func (u *UserService) Register(param model.UserRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	newUserID := uuid.New()

	user := entity.User{
		ID:       newUserID,
		Name:     param.Name,
		Email:    param.Email,
		Password: hashPassword,
		RoleID:   2,
		Level:    1,
		Xp:       0,
		Hearth:   5,
	}

	_, err = u.user.CreateUser(user)
	if err != nil {
		return err
	}

	userProfile := model.UserProfile{
		ID:     newUserID,
		Name:   param.Name,
		Email:  param.Email,
		Level:  1,
		Xp:     0,
		Hearth: 5,
	}
	if err := u.user.CreateProfile(userProfile); err != nil {
		return err
	}

	return nil
}

func (u *UserService) Login(param model.LoginAcc) (model.UserLoginResponse, error) {
	var result model.UserLoginResponse

	user, err := u.user.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}
	err = u.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwt.CreateToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (u *UserService) GetUser(param model.UserParam) (entity.User, error) {
	return u.user.GetUser(param)
}

func (u *UserService) GetProfile(userID string) (model.UserProfile, error) {
	if userID == "" {
		return model.UserProfile{}, errors.New("empty userID")
	}

	userProfile, err := u.user.SeeProfile(userID)
	if err != nil {
		return model.UserProfile{}, err
	}

	return userProfile, nil
}
