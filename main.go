package main

import (
	"os"

	"github.com/110125agnes-pixel/digital-charting-backend/database"
	"github.com/110125agnes-pixel/digital-charting-backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.DotenvLoader()
	router := gin.Default()
	database.ConnectToDigitalChartingDB()
	database.ConnectToReferenceDB()

	docMastersGroup := router.Group("/api/docMasters")
	{
		docMastersGroup.GET("/hello", func (c * gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
	}

	router.Run(":" + os.Getenv("PORT"))
}
