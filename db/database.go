package db

import (
	models "car-garage-api"
	"database/sql"
	"fmt"
	// Add GORM import and necessary initialization code
)

func GetDatabase() *gorm.DB {
	// Connect to database using credentials and Docker container connection details
	db, err := sql.Open("postgres", "...")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	// Create GORM instance and configure with db connection
	gormDB, err := gorm.Open(db)
	if err != nil {
		fmt.Println("Error initializing GORM:", err)
		return nil
	}

	// Migrate models (Car) to database tables
	gormDB.AutoMigrate(&models.Car{})

	return gormDB
}
