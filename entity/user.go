package entity

import (
	"time"

	"github.com/google/uuid"
)

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string
}

// BeforeCreate will set a UUID rather than numeric ID.
// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	user.ID = uuid.New()
// 	return
// }
