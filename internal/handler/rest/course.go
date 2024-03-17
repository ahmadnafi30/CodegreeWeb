package rest

import (
	"CodegreeWebbs/entity"
	"CodegreeWebbs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
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
