package rest

import (
	// "CodegreeWebbs/internal/service"

	// "CodegreeWebbs/entity"
	"CodegreeWebbs/entity"
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateOnboardingQuestion(ctx *gin.Context) {
	var requestData model.CreateMultipleChoiceOnboarding
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	newOnboarding := entity.Onboarding{
		Question: requestData.Question,
	}

	var options []entity.OptionBoarding
	for _, option := range requestData.Options {
		options = append(options, entity.OptionBoarding{
			Description: option.Description,
		})
	}

	newOnboarding.Options = options

	if err := r.service.OnBoardingService.SaveMultipleChoiceOnboarding(&newOnboarding, options); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to save onboarding question", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Onboarding question created successfully", nil)
}
func (r *Rest) AnswerOnBoardingQuestion(ctx *gin.Context) {
	var answer model.UserAnswerOnBoarding
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success", nil)
}

func (r *Rest) GetOnboardingQuestions(ctx *gin.Context) {
	questions, err := r.service.OnBoardingService.GetAllOnboardingQuestions()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch onboarding questions", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Onboarding questions retrieved successfully", questions)
}
