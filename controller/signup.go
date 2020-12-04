package controller

import (
	"errors"
	"graphyy/model"

	"github.com/graphql-go/graphql/gqlerrors"
)

// Signup lets users sign up for this application and returns a jwt.
func (h *BaseHandler) signup(user model.User) (model.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return model.AuthToken{}, gqlerrors.FormatError(errors.New("Invalid username"))
	}
	existingUser := h.userRepo.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return model.AuthToken{}, gqlerrors.FormatError(errors.New("this username is already in use"))
	}
	user, err := h.userRepo.SaveUser(user)
	if err != nil {
		return model.AuthToken{}, gqlerrors.FormatError(errors.New(err.Error()))
	}

	token := generateJWT(user)
	return token, nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
