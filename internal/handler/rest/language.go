package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/pkg/response"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
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

func (r *Rest) CreateMentor(ctx *gin.Context) {
	var param entity.Mentor
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed request body", err)
		return
	}
	if err := r.service.LanguageCodeService.CreateMentor(param); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to add mentor", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Added mentor successfully", nil)

}

func (r *Rest) GetAllMentor(ctx *gin.Context) {
	mentors, err := r.service.LanguageCodeService.GetAllMentor()
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "No mentors found", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Get all mentors", mentors)
}

func (r *Rest) SelectMentorWhatsAppLink(ctx *gin.Context) {
	// userIDRaw, exists := ctx.Get("userID")
	// if !exists {
	// 	response.Error(ctx, http.StatusInternalServerError, "failed to get user profile", errors.New("empty userID"))
	// 	return
	// }

	// userID, ok := userIDRaw.(uuid.UUID)
	// if !ok {
	// 	response.Error(ctx, http.StatusInternalServerError, "failed to get user profile", errors.New("userID is not a UUID"))
	// 	return
	// }

	// err := r.service.PaymentService.CheckAccess(userID)
	// if err != nil {
	// 	response.Error(ctx, http.StatusForbidden, "You don't have access to this resource", err)
	// 	return
	// }

	var requestBody struct {
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&requestBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	linkWhatsApp, err := r.service.LanguageCodeService.SelectMentorWhatsAppLink(requestBody.ID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "No mentor found", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", gin.H{"link_whatsapp": linkWhatsApp})
}
