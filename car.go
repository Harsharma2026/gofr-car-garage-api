package models

import "time"

type Car struct {
	ID           int       `gorm:"primary_key"`
	LicensePlate string    `gorm:"not null;unique"`
	EntryTime    time.Time `gorm:"not null"`
	Status       string    `gorm:"not null"`
}
