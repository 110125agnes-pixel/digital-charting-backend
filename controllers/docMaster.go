package controllers

import (
	"net/http"

	"github.com/110125agnes-pixel/digital-charting-backend/database"
	"github.com/110125agnes-pixel/digital-charting-backend/models"
	"github.com/gin-gonic/gin"
)

// GetDocMasters retrieves all records from reference.docmaster using GORM
func GetDocMasters(c *gin.Context) {
	db, ctx := database.ConnectToReferenceDB()

	var docMasters []models.DocMaster
	if err := db.WithContext(ctx).Find(&docMasters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed retrieved doc masters",
			"result":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "Successfully retrieved doc masters",
		"result":  docMasters,
	})
}
