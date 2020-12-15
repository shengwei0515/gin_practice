package models

import (
	"fmt"
	orm "seal/service/postgres"
)

type User struct {
	ID       int64  `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func (user *User) Users() (users []User, err error) {
	fmt.Print("Before query")

	result := orm.Db.Find(&users)

	if err := result.Error; err != nil {
		return nil, err
	}
	return users, err
}
