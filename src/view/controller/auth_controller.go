package controller

import (
	"ebd/src/infraestructure/middleware"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	if req.Username != "admin" || req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	userSecret := middleware.GetUserSecret(req.Username)

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"username": req.Username,
		"exp":      expirationTime.Unix(),
		"jti":      userSecret,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(userSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	parsedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	username := claims["username"].(string)

	middleware.InvalidateUserTokens(username)

	c.JSON(http.StatusOK, gin.H{"message": "Logout exitoso. Tu sesión ha sido cerrada."})
}

func SecureData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Accediste a un recurso protegido"})
}
