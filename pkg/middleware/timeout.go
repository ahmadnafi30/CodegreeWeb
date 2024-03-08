package middleware

import (
	"CodegreeWebbs/pkg/response"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func (m *middleware) Timeout() gin.HandlerFunc {
	timeLimit, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeLimit)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	response.Error(c, http.StatusRequestTimeout, "Request Time Out", errors.New(""))
}
