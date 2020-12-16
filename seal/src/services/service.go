package services

import (
	"seal/database"

	"github.com/jinzhu/gorm"
)

var DBClient *gorm.DB

func init() {
	DBClient = database.GetConnection()
}
