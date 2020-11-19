package repository

import (
	"graphyy/model"
)

// UserRepository maybe I should rename this interface
type UserRepository interface {
	GetExistingUser(username string) model.User
	SaveUser(user model.User) (model.User, error)
}
