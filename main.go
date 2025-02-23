package main

import (
	"time"

	"ebd/src/infraestructure/middleware"
	"ebd/src/view/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", controller.Login)
	r.GET("/check-db", controller.CheckDBConnection)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/secure-data", controller.SecureData)
		protected.POST("/logout", controller.Logout)
	}

	r.Run(":8585")
}
