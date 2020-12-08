package userrepo

import (
	"errors"
	"graphyy/internal"
	"graphyy/model"

	"golang.org/x/crypto/bcrypt"
)

// Login returns a jwt.
func (repo *UserRepo) Login(user model.User) (model.AuthToken, error) {
	existingUser := repo.getExistingUser(user.Username)
	if existingUser.Username == "" {
		return model.AuthToken{}, errors.New("No user found with the inputted username")
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return model.AuthToken{}, errors.New("Invalid Credentials")
	}

	token := internal.GenerateJWT(user)
	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
