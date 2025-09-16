package routes

import (
	"ecommerce/cmd/user/handler"
	"ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler, jwtSecret string) {
	//Public API
	router.Use(middleware.RequestLogger())
	router.GET("/ping", userHandler.Ping)
	router.POST("/v1/register", userHandler.RegisterRoutes)
	router.POST("/v1/login", userHandler.LoginRoutes)
	// Private API

	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	private := router.Group("/api")
	private.Use(authMiddleware)
	private.GET("/v1/userInfo", userHandler.GetUserInfo)
}
