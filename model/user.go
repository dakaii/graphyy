package model

import "time"

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Username  string     `gorm:"type:varchar(40); unique_index; not null" json:"username"`
	Password  string     `gorm:"type:varchar(40); not null" json:"password"`
}
