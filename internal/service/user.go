package service

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/bcrypt"
	"CodegreeWebbs/pkg/jwt"

	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserRegister) error
	Login(param model.LoginAcc) (model.UserLoginResponse, error)
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
