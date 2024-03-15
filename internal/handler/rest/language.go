package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/pkg/response"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateLanguage(ctx *gin.Context) {
	if r.service == nil || r.service.LanguageCodeService == nil {
		response.Error(ctx, http.StatusInternalServerError, "Internal server error", errors.New("service is not initialized"))
		return
	}

	var param entity.LanguageCode
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to parse request body", err)
		return
	}

	if err := r.service.LanguageCodeService.CreateLanguage(param); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create language", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Language created successfully", nil)
}
