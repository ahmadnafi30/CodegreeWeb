package rest

import (
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed bind inout", err)
		return
	}

	if err := r.service.UserService.Register(param); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success register new user", nil)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.LoginAcc{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login to system", token)
}
