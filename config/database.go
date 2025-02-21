package config

import (
	"example/hello/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=gogo password=myPassword dbname=mydb port=5432 sslmode=disable TimeZone=UTC"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate the User model
	if err := DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		panic("Failed to migrate database")
	}
}
