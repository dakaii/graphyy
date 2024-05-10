package userrepo

import (
	"errors"
	"fmt"
	"graphyy/entity"
	"graphyy/internal/envvar"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
func (repo *UserRepo) GetExistingUser(username string) (*entity.User, error) {
	var user *entity.User
	result := repo.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no user found with username: %s", username)
		}
		return nil, result.Error
	}
	return user, nil
}

// CreateUser creates a new user in the db..
func (repo *UserRepo) CreateUser(user entity.User) (*entity.User, error) {
	hashedPass, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass

	result := repo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("Inserted a user with ID:", user.ID)
	return &user, nil
}

func HashPassword(password string) (string, error) {
	cost := envvar.HashCost()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
