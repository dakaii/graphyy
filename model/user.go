package model

import (
	"time"

	"gorm.io/gorm"
)

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	ID        uint           `gorm:"primary_key" json:"id" graphql:"-"`
	CreatedAt time.Time      `json:"createdAt" graphql:"-"`
	UpdatedAt time.Time      `json:"updatedAt" graphql:"-"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" graphql:"-"`
	Username  string         `gorm:"unique;index; not null" json:"username" graphql:"username"`
	Password  string         `gorm:"type:varchar(1000); not null" json:"password" graphql:"password"`
}

// https://github.com/graphql-go/graphql/blob/master/examples/todo/schema/schema.go
