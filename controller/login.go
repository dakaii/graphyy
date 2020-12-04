package controller

import (
	"errors"
	"graphyy/model"

	"github.com/graphql-go/graphql/gqlerrors"
	"golang.org/x/crypto/bcrypt"
)

// Login returns a jwt.
func (h *BaseHandler) login(user model.User) (model.AuthToken, error) {
	existingUser := h.userRepo.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return model.AuthToken{}, gqlerrors.FormatError(errors.New("No user found with the inputted username"))
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return model.AuthToken{}, gqlerrors.FormatError(errors.New("Invalid Credentials"))
	}

	token := generateJWT(user)
	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
