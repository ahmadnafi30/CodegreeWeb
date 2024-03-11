package middleware

import (
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"errors"

	// "go/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AutenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, " empty token authentication", errors.New(""))
		ctx.Abort()
		return
	}
	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "Failed to validate token in authentication", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUser(model.UserParam{ID: userId})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to get user auth", err)
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
