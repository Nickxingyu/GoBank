package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
