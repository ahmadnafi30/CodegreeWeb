package jwt

import (
	// "errors"
	"CodegreeWebbs/entity"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateToken(userId uuid.UUID) (string, error)
	ValidateToken(tokenString string) (uuid.UUID, error)
	GetLoginUser(ctx *gin.Context) (entity.User, error)
}

type jsonWebToken struct {
	SecretKey   []byte
	ExpiredTime time.Duration
}

type Claims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Printf("failed set expired time for jwt: %v", err.Error())
		return nil
	}

	return &jsonWebToken{
		SecretKey:   []byte(secretKey),
		ExpiredTime: time.Duration(expiredTime) * time.Hour,
	}
}

func (j *jsonWebToken) CreateToken(userId uuid.UUID) (string, error) {
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jsonWebToken) GetLoginUser(ctx *gin.Context) (entity.User, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return entity.User{}, errors.New("Failedt get login user")
	}

	return user.(entity.User), nil

}

func (j *jsonWebToken) ValidateToken(tokenString string) (uuid.UUID, error) {
	var (
		claim  Claims
		userId uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return userId, err
	}
	if !token.Valid {
		return userId, errors.New("Invalid Token")
	}
	userId = claim.UserId
	return userId, nil
}
