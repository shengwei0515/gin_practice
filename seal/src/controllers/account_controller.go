package controllers

import (
	"net/http"
	"seal/models"
	"seal/services"

	"github.com/gin-gonic/gin"
)

// GetAccounts return all accounts in db
func GetAccounts(context *gin.Context) {

	var accounts models.Accounts
	accounts = services.GetAccounts()

	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": accounts,
	})
}
