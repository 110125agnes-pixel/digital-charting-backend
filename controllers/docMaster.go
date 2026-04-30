package controllers

import (
	"net/http"

	"github.com/110125agnes-pixel/digital-charting-backend/database"
	"github.com/110125agnes-pixel/digital-charting-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func GetDocMasterByDocCode(c * gin.Context) {
	docCode := c.Param("docCode");
	db, ctx := database.ConnectToReferenceDB()

	docMaster, err := gorm.G[models.DocMaster](db).Where("doccode = ?", docCode).First(ctx)

	if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"message": "Failed to retrieve doc master",
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result" : docMaster,
		"message" : "Get docMaster successfully",
		"error": nil,
	})
}