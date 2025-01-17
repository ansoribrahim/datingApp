package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"datingApp/models"
	"datingApp/services"
)

func RegisterAuthRoutes(router *gin.Engine, authService *services.AuthService) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", func(c *gin.Context) {
			var userRequest models.SignUpRequest
			if err := c.ShouldBindJSON(&userRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			resp, err := authService.SignUp(userRequest)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, resp)
		})

		authGroup.POST("/login", func(c *gin.Context) {
			var loginReq models.LoginRequest

			if err := c.ShouldBindJSON(&loginReq); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			user, err := authService.Login(loginReq.Email, loginReq.Password)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"user": user})
		})
	}
}
