package main

import (
	"github.com/110125agnes-pixel/digital-charting-backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ping := router.Group("/api")
	{
		ping.GET("/ping", controllers.Ping)
	}

	router.Run(":8000")
}