package controller

import (
	"graphyy/controller/user"
	"graphyy/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	userController *user.UserController
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		userController: user.NewUserController(repositories.UserRepo),
	}
}
