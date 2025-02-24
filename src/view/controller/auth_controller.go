package controller

import (
	"ebd/src/infraestructure/middleware"
	"ebd/src/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	loginUseCase := auth.LoginUseCase{}

	userDTO, err := loginUseCase.Execute(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Respuesta de autenticación exitosa con UserDTO
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    userDTO,
	})
}

func Logout(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}

	middleware.InvalidateUserTokens(string(userID.(int64)))

	c.JSON(http.StatusOK, gin.H{"message": "Logout exitoso. Tu sesión ha sido cerrada."})
}

func SecureData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Accediste a un recurso protegido"})
}
