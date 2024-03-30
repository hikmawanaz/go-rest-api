package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=xbie password= dbname=belyfe port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Check the connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	if err := sqlDB.Ping(); err != nil {
		panic("failed to ping database")
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
}