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

	ping := router.Group("/api")
	{
		ping.GET("/ping", controllers.Ping)
	}

	router.Run(":" + os.Getenv("PORT"))
}
