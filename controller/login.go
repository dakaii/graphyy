package controller

import (
	"errors"
	"graphyy/model"

	"golang.org/x/crypto/bcrypt"
)

// Login returns a jwt.
func (h *BaseHandler) Login(user model.User) (*model.AuthToken, error) {
	existingUser := h.userRepo.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return nil, errors.New("No user found with the inputted username")
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return nil, errors.New("Invalid Credentials")
	}

	token := generateJWT(user)
	return &token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
