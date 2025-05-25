package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"-" gorm:"default:null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
}