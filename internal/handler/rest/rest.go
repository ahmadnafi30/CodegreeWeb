package rest

import (
	"CodegreeWebbs/internal/service"
	"CodegreeWebbs/pkg/middleware"
	"CodegreeWebbs/pkg/response"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoints() {
	r.router.Use(r.middleware.Timeout())
	routerGroup := r.router.Group("/api/v1/")
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/profile", r.GetProfile)
	routerGroup.GET("/login-user", r.middleware.AutenticateUser, r.GetLoginUser)

	routerGroup.POST("/create_onboarding_question", r.CreateOnboardingQuestion)
	routerGroup.GET("/onboarding_questions", r.GetOnboardingQuestions)
	routerGroup.POST("/answer_boarding", r.AnswerOnBoardingQuestion)
}

func (r *Rest) Run() {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	if addr == "" {
		addr = "127.0.0.1"
	}

	if port == "" {
		port = "8080"
	}

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func (r *Rest) TestTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	response.Success(ctx, http.StatusOK, "success", nil)
}
