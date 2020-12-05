package userrepo

import (
	"errors"
	"graphyy/internal"
	"graphyy/model"
)

// Signup lets users sign up for this application and returns a jwt.
func (repo *UserRepo) Signup(user model.User) (model.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return model.AuthToken{}, errors.New("Invalid username")
	}
	existingUser := repo.getExistingUser(user.Username)
	if existingUser.Username != "" {
		return model.AuthToken{}, errors.New("this username is already in use")
	}
	user, err := repo.saveUser(user)
	if err != nil {
		return model.AuthToken{}, err
	}

	token := internal.GenerateJWT(user)
	return token, nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
