package routers

import (
	"seal/controllers"

	"github.com/gin-gonic/gin"
)

// AddRoutesAccounts return RouterGroup for Acccount CRUD
func AddRoutesAccounts(r *gin.RouterGroup, controller *controllers.AccountController) {
	r.GET("/accounts", controller.GetAccounts)
}
