package routers

import (
	"github.com/gin-gonic/gin"
)

// InitiRouter

// a function return pointer of gin.Engine
func InitRouters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routerGroup := router.Group("api")
	routerGroup.Use()
	{
		AddRoutesAccounts(routerGroup)
	}

	return router
}
