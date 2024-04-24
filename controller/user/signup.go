package user

import (
	"errors"
	"graphyy/entity"
	"graphyy/internal"
)

// Signup lets users sign up for this application and returns a jwt.
func (c *Controller) Signup(user entity.User) (entity.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return entity.AuthToken{}, errors.New("Invalid username")
	}
	existingUser := c.service.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return entity.AuthToken{}, errors.New("this username is already in use")
	}
	user, err := c.service.CreateUser(user)
	if err != nil {
		return entity.AuthToken{}, err
	}

	token := internal.GenerateJWT(user)
	return token, nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
