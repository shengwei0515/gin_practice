package router

import (
	. "seal/controller"

	"github.com/gin-gonic/gin"
)

// InitiRouter

// a function return pointer of gin.Engine
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", GetUsers)
	return router
}
