package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	*gin.Engine
}

func (r AccountRouter) Load() {
	rg := r.Group("/account")

	rg.GET("/:id", getAccountById)
	rg.POST("/", createAccount)
	rg.PUT("/", updateAccount)
	rg.DELETE("/:id", deleteAccountById)
}

func getAccountById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Account By Id",
	})
}

func createAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create Account",
	})
}

func updateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Account",
	})
}

func deleteAccountById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Account By Id",
	})
}
