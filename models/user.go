package models

// TODO probably not the best package name. check what the best practice is.

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetExistingUser(username string) User
	SaveUser(user User) (User, error)
}
