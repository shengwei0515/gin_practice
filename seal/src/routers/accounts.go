package routers

import (
	controllers "seal/controllers"

	"github.com/gin-gonic/gin"
)

// AddRoutesAccounts return RouterGroup for Acccount CRUD
func AddRoutesAccounts(r *gin.RouterGroup) {
	r.GET("/accounts", controllers.GetAccounts)
}
