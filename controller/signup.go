package controller

import (
	"errors"
	"graphyy/model"
)

// Signup lets users sign up for this application and returns a jwt.
func (h *BaseHandler) signup(user model.User) (model.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return model.AuthToken{}, errors.New("Invalid username")
	}
	existingUser := h.userRepo.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return model.AuthToken{}, errors.New("this username is already in use")
	}
	user, _ = h.userRepo.SaveUser(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	token := generateJWT(user)
	return token, nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
