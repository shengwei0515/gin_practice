package services

import (
	"seal/models"
)

// GetAccounts return all accounts
func GetAccounts() models.Accounts {

	var accounts models.Accounts
	dbClient := DBClient

	result := dbClient.Table("Accounts").Take(&accounts)
	if err := result.Error; err != nil {
		return nil
	}

	return accounts
}
