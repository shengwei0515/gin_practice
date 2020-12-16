package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const (
	ENV_DB_DRIVER   = "type"
	ENV_DB_HOST     = "host"
	ENV_DB_NAME     = "name"
	ENV_DB_USER     = "user"
	ENV_DB_PASSWORD = "password"
)

var (
	DB_DRIVER   string = "postgres"
	DB_HOST     string = "127.0.0.1"
	DB_NAME     string = "postgres"
	DB_USER     string = "postgres"
	DB_PASSWORD string = "password"
)

var Db *gorm.DB
var DbErr error

func init() {
	getEnvDatabaseConfig()
}

func getEnvDatabaseConfig() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("database.env")
	if err != nil {
		fmt.Print(err)
	}

	dbType := myEnv["type"]
	dbUser := myEnv["user"]
	dbPwd := myEnv["password"]
	dbName := myEnv["name"]
	dbHost := myEnv["host"]

	// value handel
	if len(dbType) > 0 {
		DB_DRIVER = dbType
	}
	if len(dbUser) > 0 {
		DB_USER = dbUser
	}
	if len(dbPwd) > 0 {
		DB_PASSWORD = dbPwd
	}
	if len(dbName) > 0 {
		DB_NAME = dbName
	}
	if len(dbHost) > 0 {
		DB_HOST = dbHost
	}
}

// GetConnection return postgres db connection
func GetConnection() *gorm.DB {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, DB_USER, DB_NAME, DB_PASSWORD)
	db, err := gorm.Open(DB_DRIVER, dbURI)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
