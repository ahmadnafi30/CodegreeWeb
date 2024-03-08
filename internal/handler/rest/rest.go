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

func (r *Rest) MountEndPoin() {
	r.router.Use(r.middleware.Timeout())
	routerGroup := r.router.Group("/api/v1/")
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)

}

func (r *Rest) Run() {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)
	response.Success(ctx, http.StatusOK, "succes", nil)
}
