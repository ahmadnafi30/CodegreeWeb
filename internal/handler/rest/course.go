package rest

import (
	"CodegreeWebbs/model"
	"CodegreeWebbs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateCourse(ctx *gin.Context) {
	var course model.CreateCourse
	if err := ctx.ShouldBindJSON(&course); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid request", err)
		return
	}
}
