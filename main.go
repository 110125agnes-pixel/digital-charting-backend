package main

import (
	"os"

	"github.com/110125agnes-pixel/digital-charting-backend/controllers"
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
		docMastersGroup.GET("/", controllers.GetDocMasters)
		docMastersGroup.GET("/:docCode", controllers.GetDocMasterByDocCode)
	}

	router.Run(":" + os.Getenv("PORT"))
}
