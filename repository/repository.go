package repository

import (
	"graphyy/repository/userrepo"

	"gorm.io/gorm"
)

// Repositories contains all the repo structs
type Repositories struct {
	UserRepo *userrepo.UserRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepo(db)
	return &Repositories{UserRepo: userRepo}
}
