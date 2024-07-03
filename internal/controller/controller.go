package controller

import (
	"github.com/dakaii/graphyy/internal/controller/user"
	"github.com/dakaii/graphyy/internal/repository"
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
