package user

import (
	"errors"
	"graphyy/entity"
	"graphyy/internal"

	"golang.org/x/crypto/bcrypt"
)

// Login returns a jwt.
func (c *Controller) Login(user entity.User) (entity.AuthToken, error) {
	existingUser := c.service.GetExistingUser(user.Username)
	if existingUser.Username == "" {
		return entity.AuthToken{}, errors.New("No user found with the inputted username")
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return entity.AuthToken{}, errors.New("Invalid Credentials")
	}

	token := internal.GenerateJWT(user)
	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
