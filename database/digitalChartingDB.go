package database

import (
	"context"
	"os"

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
	
	// *** put all db.AutoMigrate() between this comment ***

	ctx := context.Background()

	return db, ctx
}

