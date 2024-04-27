package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;index;not null"`
	Password  string         `gorm:"type:varchar(1000);not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
