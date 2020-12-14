package models

type User struct {
	ID       int64  `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func (user *User) Users() (users []User, err error) {
	// TODO: add db query and return users

	userFound := []User{}

	return userFound, nil
}
