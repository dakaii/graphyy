package repository

import (
	"fmt"
	"graphyy/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository maybe I should rename this interface
type UserRepository interface {
	GetExistingUser(username string) model.User
	SaveUser(user model.User) (model.User, error)
}

// UserRepo should i rename it?
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (h *UserRepo) GetExistingUser(username string) model.User {
	var user model.User
	h.db.Where("username = ?", username).First(*&user)
	return model.User{Username: user.Username, Password: user.Password}
}

// SaveUser creates a new user in the db..
func (h *UserRepo) SaveUser(user model.User) (model.User, error) {
	// TODO handle the potential error below.
	hashedPass, _ := hashPassword(user.Password)
	user.Password = hashedPass

	h.db.Create(&user)
	// result := h.db.Create(&user)
	// if result.Error
	fmt.Println("Inserted a user with ID:", user.ID)
	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
