package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/pkg/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateCourse(ctx *gin.Context) {
	var courseData entity.Course
	if err := ctx.BindJSON(&courseData); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	if err := r.service.CourseService.CreateCourse(&courseData); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create course", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Create Course success", nil)
}

func (r *Rest) GetAllCourses(ctx *gin.Context) {
	courses, err := r.service.CourseService.GetAllCourses()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get courses", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Get all courses success", courses)
}

func (r *Rest) SelectCourse(ctx *gin.Context) {

	var requestBody struct {
		ID uint `json:"id"`
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to parse request body", err)
		return
	}
	if requestBody.ID == 0 {
		response.Error(ctx, http.StatusBadRequest, "Invalid course ID", errors.New("empty course ID"))
		return
	}

	courseDetail, err := r.service.CourseService.Selectacourse(requestBody.ID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get the course", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Course retrieved successfully", courseDetail)
}

func (r *Rest) SelectSubLang(ctx *gin.Context) {
	var requestBody struct {
		ID uint `json:"id"`
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to parse request body", err)

	}
	if requestBody.ID == 0 {
		response.Error(ctx, http.StatusBadRequest, "Invalid sublang ID", errors.New("empty course ID"))

	}
	sublang, err := r.service.CourseService.SelectSubLang(requestBody.ID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to select the sublang", err)

	}
	response.Success(ctx, http.StatusInternalServerError, "Selected the sublang successfully.", sublang)
}

func (r *Rest) GetGamification(ctx *gin.Context) {
	var requestBody struct {
		SublangID uint
		ID        uint `json:"id"`
	}
	if err := ctx.BindJSON(&requestBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Dailed to parse rquest body", err)
	}
	if requestBody.ID == 0 {
		response.Error(ctx, http.StatusBadRequest, "Invalid question ID", errors.New("empty course ID"))

	}
	question, err := r.service.CourseService.GetGamification(requestBody.SublangID, requestBody.ID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get question", err)

	}
	response.Success(ctx, http.StatusInternalServerError, "Success get question", question)

}

func (r *Rest) CheckAnswer(ctx *gin.Context) {

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

	var request struct {
		Question uint `json:"question"`
		Option   uint `json:"option"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Bad Request", err)
		return
	}

	correctAnswer, err := r.service.CourseService.CheckAnswer(userID, request.Question, request.Option)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	if correctAnswer {
		response.Success(ctx, http.StatusOK, "Correct answer", nil)
	} else {
		response.Success(ctx, http.StatusOK, "Incorrect answer", nil)
	}
}
