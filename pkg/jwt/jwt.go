package jwt

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Interface interface {
	CreateToken(userId uuid.UUID) (string, error)
}

type jsonWebToken struct {
	SecretKey   string
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
		log.Fatalf("failed set expired time for jwt : %v", err.Error())
	}

	return &jsonWebToken{
		SecretKey:   secretKey,
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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(j.SecretKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
