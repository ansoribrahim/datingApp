package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"datingApp/middleware"
	"datingApp/models"
	"datingApp/services"
)

func RegisterSwipeRoutes(router *gin.Engine, swipeService *services.SwipeService, jwtSecret string) {
	swipeGroup := router.Group("/swipe")
	// Apply JWTAuth middleware
	swipeGroup.Use(middleware.JWTAuth(jwtSecret))
	{
		swipeGroup.POST("/right", func(c *gin.Context) {
			// Extract user ID from JWT claims
			userIDStr, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			// Convert userID string to uuid.UUID
			userID, err := uuid.Parse(userIDStr.(string))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
				return
			}

			var req models.SwipeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			err = swipeService.SwipeRight(userID, req.ProfileID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Swipe right recorded"})
		})

		swipeGroup.POST("/left", func(c *gin.Context) {
			// Extract user ID from JWT claims
			userIDStr, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			// Convert userID string to uuid.UUID
			userID, err := uuid.Parse(userIDStr.(string))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
				return
			}

			var req models.SwipeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}

			err = swipeService.SwipeLeft(userID, req.ProfileID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Swipe left recorded"})
		})

		swipeGroup.GET("/matches", func(c *gin.Context) {
			// Extract user ID from JWT claims
			userIDStr, exists := c.Get("userID")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			// Convert userID string to uuid.UUID
			userID, err := uuid.Parse(userIDStr.(string))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
				return
			}

			matches, err := swipeService.GetPotentialMatches(userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"matches": matches})
		})
	}
}
