package db

import (
	models "car-garage-api"
	"database/sql"
	"fmt"
)

func GetDatabase() *gorm.DB {

	db, err := sql.Open("postgres", "...")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	gormDB, err := gorm.Open(db)
	if err != nil {
		fmt.Println("Error initializing GORM:", err)
		return nil
	}

	gormDB.AutoMigrate(&models.Car{})

	return gormDB
}
