package controllers

import (
	"net/http"
	"seal/models"
	"seal/services"

	"github.com/gin-gonic/gin"
)

// IAccountController is a list of function for AccountController
type IAccountController interface {
	GetAccount(context *gin.Context)
}

// AccountController is a controller to handal account CRUD
type AccountController struct {
	AccountService *services.AccountService
}

// GetAccounts return all accounts in db
func (accountcontroller *AccountController) GetAccounts(context *gin.Context) {

	var accounts models.Accounts
	accounts = accountcontroller.AccountService.GetAccounts()

	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": accounts,
	})
}
