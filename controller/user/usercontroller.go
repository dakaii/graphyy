package user

import (
	"graphyy/repository/userrepo"
)

type UserController struct {
	userRepository userrepo.UserRepository
}

// // NewBaseHandler returns a new BaseHandler
func NewUserController(userRepo *userrepo.UserRepo) *UserController {
	return &UserController{
		userRepository: userRepo,
	}
}
