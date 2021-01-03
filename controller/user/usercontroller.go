package user

import (
	"graphyy/model"
	"graphyy/repository/userrepo"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type repository interface {
	GetExistingUser(username string) model.User
	CreateUser(user model.User) (model.User, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	service repository
}

// InitController initializes the user controller.
func InitController(userRepo *userrepo.UserRepo) *Controller {
	return &Controller{
		service: userRepo,
	}
}
