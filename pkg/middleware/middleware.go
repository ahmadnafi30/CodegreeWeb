package middleware

import (
	"CodegreeWebbs/internal/service"
	"CodegreeWebbs/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Timeout() gin.HandlerFunc
	AuthenticateUser(ctx *gin.Context)
	OnlySubscription(ctx *gin.Context)
}

type middleware struct {
	jwtAuth jwt.Interface
	service *service.Service
}

func Init(jwtAuth jwt.Interface, service *service.Service) Interface {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}
