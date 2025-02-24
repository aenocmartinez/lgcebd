package main

import (
	"time"

	"ebd/src/infraestructure/middleware"
	"ebd/src/view/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Configuración de CORS
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

		cursoGroup := protected.Group("/cursos")
		{
			cursoGroup.GET("/", controller.ListarCursos)
			cursoGroup.POST("/", controller.CrearCurso)
			cursoGroup.PUT("/:id", controller.ActualizarCurso)
		}
	}

	r.Run(":8585")
}
