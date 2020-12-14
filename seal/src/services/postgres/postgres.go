package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var Db *gorm.DB

func init() {

	// Load .evn file to get config
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

	fmt.Println(dbType)
	// gen connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPwd)
	fmt.Println(dbURI)

	// connect to db
	Db, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}
	if Db.Error != nil {
		fmt.Print(Db.Error)
	}
}
