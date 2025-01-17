package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"datingApp/models"
	"datingApp/services"
)

func RegisterPremiumRoutes(r *gin.Engine, premiumService *services.PremiumService) {
	premium := r.Group("/premium")
	{
		// Get all premium packages
		premium.GET("/packages", func(c *gin.Context) {
			packages, err := premiumService.GetAllPremiumPackages()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch premium packages"})
				return
			}
			c.JSON(http.StatusOK, packages)
		})

		// Get specific premium package
		premium.GET("/packages/:packageID", func(c *gin.Context) {
			packageID, err := uuid.Parse(c.Param("packageID"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID format"})
				return
			}

			pkg, err := premiumService.GetPremiumPackageByID(packageID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch premium package"})
				return
			}
			if pkg == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
				return
			}
			c.JSON(http.StatusOK, pkg)
		})

		// Create new premium package (admin only)
		premium.POST("/packages", func(c *gin.Context) {
			var req models.PremiumPackageRequest // Assuming this is defined in `models`

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			price, err := decimal.NewFromString(req.Price)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
				return
			}

			err = premiumService.CreatePremiumPackage(req.Name, req.Description, price)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create premium package"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "Premium package created successfully"})
		})

		// Purchase premium package
		premium.POST("/purchase", func(c *gin.Context) {
			var req models.PurchaseRequest // Assuming this is defined in `models`

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			userID, err := uuid.Parse(req.UserID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
				return
			}

			packageID, err := uuid.Parse(req.PackageID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID format"})
				return
			}

			if err := premiumService.RegisterPremium(userID, packageID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Premium package purchased successfully"})
		})

		// Check premium status
		premium.GET("/status/:userID", func(c *gin.Context) {
			userID, err := uuid.Parse(c.Param("userID"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
				return
			}

			isPremium, err := premiumService.IsUserPremium(userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check premium status"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"is_premium": isPremium})
		})

		// Update premium package (admin only)
		premium.PUT("/packages/:packageID", func(c *gin.Context) {
			packageID, err := uuid.Parse(c.Param("packageID"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID format"})
				return
			}

			var req models.PremiumPackageRequest // Assuming this is defined in `models`

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			price, err := decimal.NewFromString(req.Price)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price format"})
				return
			}

			pkg := &models.PremiumPackage{
				ID:          packageID,
				PackageName: req.Name,
				Description: req.Description,
				Price:       price,
			}

			if err := premiumService.UpdatePremiumPackage(pkg); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update premium package"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Premium package updated successfully"})
		})

		// Delete premium package (admin only)
		premium.DELETE("/packages/:packageID", func(c *gin.Context) {
			packageID, err := uuid.Parse(c.Param("packageID"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID format"})
				return
			}

			if err := premiumService.DeletePremiumPackage(packageID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete premium package"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Premium package deleted successfully"})
		})
	}
}
