package entity

import (
	"time"

	"gorm.io/gorm"
)

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Username  string         `gorm:"unique;index;not null" json:"username"`
	Password  string         `gorm:"type:varchar(1000);not null" json:"password"`
}
