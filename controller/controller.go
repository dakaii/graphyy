package controller

import (
	"graphyy/controller/user"
	"graphyy/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	UserController *user.Controller
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		UserController: user.InitController(repositories.UserRepo),
	}
}
