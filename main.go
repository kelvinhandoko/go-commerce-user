package main

import (
	"ecommerce/cmd/user/handler"
	"ecommerce/cmd/user/repository"
	"ecommerce/cmd/user/resource"
	"ecommerce/cmd/user/service"
	"ecommerce/cmd/user/usecase"
	"ecommerce/config"
	"ecommerce/infrastructure/log"
	"ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.IntDb(&cfg)
	log.SetupLoger()

	userRepository := repository.NewUserRepository(redis, db)
	userService := service.NewUserService(*userRepository)
	userUseCase := usecase.NewUserUseCase(*userService, cfg.Secret.JWTSecret)
	UserHandler := handler.NewUserHandler(*userUseCase)

	port := cfg.App.Port
	router := gin.Default()

	routes.SetupRoutes(router, *UserHandler, cfg.Secret.JWTSecret)

	router.Run(":" + port)

	log.Logger.Printf("Server running on port %s", port)

}
