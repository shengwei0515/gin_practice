package models

// Account is a single account with id ,username and password
type Account struct {
	ID       int64  `json: "id"`
	Username string `json: "username" binding:"required"`
	Password string `json: "password" binding:"required"`
}

// Accounts is a list of account
type Accounts []Account
