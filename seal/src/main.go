package main

import (
	"seal/controllers"
	"seal/database"
	"seal/routers"
	"seal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// get db connection
	postgresDB := database.GetConnection()
	// new services
	accountService := &services.AccountService{
		DBClient: postgresDB,
	}

	// new controllers
	accountController := &controllers.AccountController{
		AccountService: accountService,
	}
	// setup routing and run
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routerGroup := router.Group("api")
	routerGroup.Use()
	routers.AddRoutesAccounts(routerGroup, accountController)

	router.Run(":8080")
}
