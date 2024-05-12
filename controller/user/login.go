package user

import (
	"errors"
	"graphyy/domain"

	"graphyy/internal/auth"

	"golang.org/x/crypto/bcrypt"
)

// Login returns a jwt.
func (c *Controller) Login(user domain.User) (domain.AuthToken, error) {
	existingUser, err := c.service.GetExistingUser(user.Username)
	if err != nil {
		return domain.AuthToken{}, errors.New("no user found with the inputted username")
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return domain.AuthToken{}, errors.New("invalid credentials")
	}

	token := auth.GenerateJWT(user)
	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
