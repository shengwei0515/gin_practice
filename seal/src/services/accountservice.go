package services

import (
	"errors"
	"seal/models"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

// IAccountService interface
type IAccountService interface {
	GetAccounts() models.Accounts
	CreateAccount(context *gin.Context) (int64, error)
}

// AccountService struct
type AccountService struct {
	DBClient *gorm.DB
}

// GetAccounts return all accounts
func (accountSvc *AccountService) GetAccounts() models.Accounts {

	var accounts models.Accounts

	result := accountSvc.DBClient.Table("Accounts").Take(&accounts)
	if err := result.Error; err != nil {
		return nil
	}

	return accounts
}

func (accountSvc *AccountService) CreateAccount(context *gin.Context) (int64, error){
	var account models.Account
	context.BindJSON(&account)
	if _ValidateAccountColumn(&account) {
		result := accountSvc.DBClient.Table("Accounts").Create(&account)
		if result.Error == nil {
			return result.RowsAffected, nil
		} else {
			return 0, result.Error
		}
	} else{
		return 0, errors.New("username and password can not be nil")
	}
}

func _ValidateAccountColumn(account *models.Account) bool {
	if account.Username == "" || account.Password == ""{
		return false
	} else {
		return true
	}
}