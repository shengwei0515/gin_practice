package services

import (
	"seal/models"

	"github.com/jinzhu/gorm"
)

// IAccountService interface
type IAccountService interface {
	GetAccounts() models.Accounts
}

// AccountService struct
type AccountService struct {
	DBClient *gorm.DB
}

// GetAccounts return all accounts
func (accountsvc *AccountService) GetAccounts() models.Accounts {

	var accounts models.Accounts

	result := accountsvc.DBClient.Table("Accounts").Take(&accounts)
	if err := result.Error; err != nil {
		return nil
	}

	return accounts
}
