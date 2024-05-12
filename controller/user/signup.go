package user

import (
	"errors"
	"graphyy/domain"
	"graphyy/internal/auth"
)

// Signup lets users sign up for this application and returns a jwt.
func (c *Controller) Signup(user domain.User) (domain.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return domain.AuthToken{}, errors.New("invalid username")
	}
	existingUser, _ := c.service.GetExistingUser(user.Username)
	if existingUser != nil {
		return domain.AuthToken{}, errors.New("this username is already in use")
	}

	createdUser, err := c.service.CreateUser(user)
	if err != nil {
		return domain.AuthToken{}, err
	}

	token := auth.GenerateJWT(*createdUser)
	return token, nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
