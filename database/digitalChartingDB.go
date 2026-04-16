package database

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDigitalChartingDB() (*gorm.DB, context.Context) {
	dsn := "root:root@tcp(localhost:3306)/digital-charting?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed to connect to database")
	}

	ctx := context.Background()

	return db, ctx
}

