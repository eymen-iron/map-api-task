package models

import "gorm.io/gorm"

type Location struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	Name      string  `gorm:"type:varchar(45);not null"`
	Latitude  float64 `gorm:"type:float;not null"`
	Longitude float64 `gorm:"type:float;not null"`
	Marker    string  `gorm:"type:varchar(45);not null"`
	Distance  float64 `gorm:"-"`
	gorm.Model
}
