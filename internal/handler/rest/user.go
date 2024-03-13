package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"errors"

	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed bind input", err)
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

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid email or password", err)
		return
	}

	token, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success login to system", token)
}

func (r *Rest) GetLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "failed get login user", errors.New("not found login user"))
		return
	}
	response.Success(ctx, http.StatusOK, "Success get user", user.(entity.User))
}

func (r *Rest) GetProfile(ctx *gin.Context) {
	userIDRaw, exists := ctx.Get("userID")
	if !exists {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user profile", errors.New("empty userID"))
		return
	}

	userID, ok := userIDRaw.(uuid.UUID)
	if !ok {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user profile", errors.New("userID is not a UUID"))
		return
	}

	userIDStr := userID.String()

	userProfile, err := r.service.UserService.GetProfile(userIDStr)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user profile", err)
		return
	}
	response.Success(ctx, http.StatusOK, "success get user profile", userProfile)
}
