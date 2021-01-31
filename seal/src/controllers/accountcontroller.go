package controllers

import (
	"fmt"
	"net/http"
	"seal/models"
	"seal/services"
	"github.com/gin-gonic/gin"
)

// IAccountController is a list of function for AccountController
type IAccountController interface {
	GetAccount(context *gin.Context)
	CreateAccount(context *gin.Context)
}

// AccountController is a controller to handal account CRUD
type AccountController struct {
	AccountService *services.AccountService
}

// GetAccounts return all accounts in db
func (accountController *AccountController) GetAccounts(context *gin.Context) {

	var accounts models.Accounts
	accounts = accountController.AccountService.GetAccounts()

	context.JSON(http.StatusOK, gin.H{
		"code": len(accounts),
		"data": accounts,
	})
}

func (accountController * AccountController) CreateAccount(context *gin.Context) {
	rowAffected, err := accountController.AccountService.CreateAccount(context)
	
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"rowAffected": nil,
			"message": "Error Create",
		})
		fmt.Println(err)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"rowAffected": rowAffected,
			"message": "Create Successed",
		})
	}
}
