package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateOnboardingQuestion(ctx *gin.Context) {
	var requestData model.CreateMultipleChoiceOnboarding
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
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
		response.Error(ctx, http.StatusInternalServerError, "saving  onboarding question failed", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to saving question", nil)
}

func (r *Rest) GetOnboardingQuestions(ctx *gin.Context) {
	questions, err := r.service.OnBoardingService.GetAllOnboardingQuestions()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch onboarding questions", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Onboarding questions retrieved successfully", questions)
}

func (r *Rest) AnswerOnboardingQuestion(ctx *gin.Context) {
	userIDRaw, _ := ctx.Get("userID")
	userID, _ := userIDRaw.(uuid.UUID)

	_, err := r.service.OnBoardingService.GetUserByID(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user id", err)
		return
	}

	var answer entity.UserAnswerOnBoarding
	if err := ctx.ShouldBindJSON(&answer); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Request invalid", err)
		return
	}

	if answer.QuestionID == 0 || answer.Answer == 0 {
		response.Error(ctx, http.StatusBadRequest, "id or answer invalid", nil)
		return
	}

	answer.UserID = userID

	if err := r.service.OnBoardingService.SaveUserAnswer(&answer); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to save to database", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success save to database", nil)
}

func (r *Rest) RecommendLanguage(ctx *gin.Context) {
	userIDRaw, _ := ctx.Get("userID")
	userID, _ := userIDRaw.(uuid.UUID)

	recommendedLanguages, err := r.service.OnBoardingService.CheckAnswer(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get recommended languages", err)
		return
	}
	if len(recommendedLanguages) == 0 {
		response.Error(ctx, http.StatusNotFound, "No recommended languages found", nil)
		return
	}
	response.Success(ctx, http.StatusOK, "Recommended languages found", recommendedLanguages)
}
