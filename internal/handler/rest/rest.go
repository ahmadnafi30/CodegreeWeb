package rest

import (
	"CodegreeWebbs/internal/service"
	"CodegreeWebbs/pkg/middleware"
	"CodegreeWebbs/pkg/response"
	"fmt"

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
	r.router.Use(middleware.Cors())
	routerGroup := r.router.Group("/api/v1")
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/profile", r.middleware.AuthenticateUser, r.GetProfile)

	routerGroup.POST("/create_onboarding_question", r.CreateOnboardingQuestion)
	routerGroup.GET("/onboarding_questions", r.middleware.AuthenticateUser, r.GetOnboardingQuestions)
	routerGroup.POST("/answer_onboarding_question", r.middleware.AuthenticateUser, r.AnswerOnboardingQuestion)
	routerGroup.GET("/recomend_language", r.middleware.AuthenticateUser, r.RecommendLanguage)

	// routerGroup.POST("/create_language", r.CreateLanguage)
	routerGroup.POST("/create_course", r.CreateCourse)
	routerGroup.GET("/get_course", r.GetAllCourses)
	routerGroup.POST("/select_course", r.SelectCourse)
	routerGroup.POST("/select_sublang", r.SelectSubLang)
	routerGroup.POST("/select_question", r.GetGamification)
	routerGroup.POST("/answer_quest", r.middleware.AuthenticateUser, r.CheckAnswer)
	routerGroup.POST("/addmentor", r.CreateMentor)
	routerGroup.GET("/get_mentors", r.GetAllMentor)
}

func (r *Rest) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r.router.Run(fmt.Sprintf(":%s", port))
}

func (r *Rest) TestTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	response.Success(ctx, http.StatusOK, "success", nil)
}
