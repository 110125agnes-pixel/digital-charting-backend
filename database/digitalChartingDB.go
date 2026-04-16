package database

import (
	"context"
	"os"

	"github.com/110125agnes-pixel/digital-charting-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDigitalChartingDB() (*gorm.DB, context.Context) {
	dsn := os.Getenv("DIGITAL_CHARTING_DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed to connect to database")
	}

	// *** put all db.AutoMigrate() between this comment ***
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	// *** put all db.AutoMigrate() between this comment ***

	ctx := context.Background()

	return db, ctx
}

