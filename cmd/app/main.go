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

	middleware := middleware.Init()

	r := rest.NewRest(svc, middleware)

	r.MountEndPoin()
	r.Run()
}
