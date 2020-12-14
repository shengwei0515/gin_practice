package controller

import (
	"fmt"
	"net/http"
	model "seal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	var user model.User
	result, err := user.Users()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Can not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
