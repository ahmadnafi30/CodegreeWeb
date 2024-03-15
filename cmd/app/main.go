package main

import (
	"CodegreeWebbs/internal/handler/rest"
	"CodegreeWebbs/internal/repository"
	"CodegreeWebbs/internal/service"
	"CodegreeWebbs/pkg/bcrypt"
	"CodegreeWebbs/pkg/config"
	"CodegreeWebbs/pkg/database/mysql"
	"CodegreeWebbs/pkg/jwt"
	"CodegreeWebbs/pkg/middleware"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := mysql.ConnectDatabase()
	if err := mysql.Migration(db); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)

	bcryptService := bcrypt.Init()
	jwtAuth := jwt.Init()

	svc := service.NewService(service.InitParam{
		Repository: repo,
		Bcrypt:     bcryptService,
		JwtAuth:    jwtAuth,
	})
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	router.Run()

	middleware := middleware.Init(jwtAuth, svc)

	r := rest.NewRest(svc, middleware)

	r.MountEndpoints()
	r.Run()

}
