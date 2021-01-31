package routers

import (
	"seal/controllers"

	"github.com/gin-gonic/gin"
)

// AddRoutesAccount return RouterGroup for Acccount CRUD
func AddRoutesAccount(r *gin.RouterGroup, controller *controllers.AccountController) {
	r.GET("/accounts", controller.GetAccounts)
	r.POST("/account", controller.CreateAccount)
}
