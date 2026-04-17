package database

import (
	"context"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToReferenceDB() (*gorm.DB, context.Context) {
	dsn := os.Getenv("REFERENCE_DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to reference database")
	}
	ctx := context.Background()
	return db, ctx
}
