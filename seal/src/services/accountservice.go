package services

import (
	"errors"
	"seal/models"

	"github.com/jinzhu/gorm"
)

// IAccountService interface
type IAccountService interface {
	GetAccounts() models.Accounts
	CreateAccount(account *models.Account) (int64, error)
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

func (accountSvc *AccountService) CreateAccount(account *models.Account) (int64, error){
	if _ValidateAccountColumn(account) {
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

func (accountSvc *AccountService) UpdateAccount(account *models.Account) (int64, error){
	if _ValidateAccountColumn(account) {
		result := accountSvc.DBClient.Table("Accounts").Where("username = ?",account.Username).Update("password", account.Password)
		if result.Error == nil && result.RowsAffected != 0{
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