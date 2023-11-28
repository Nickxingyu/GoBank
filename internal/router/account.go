package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	*gin.Engine
}

func (r AccountRouter) load() {
	rg := r.Group("/account")

	rg.GET("/:id", getAccountById)
	rg.POST("/", createAccount)
	rg.PUT("/", updateAccount)
	rg.DELETE("/:id", deleteAccountById)
}

func getAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get An Account By ID : " + id,
	})
}

func createAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create An Account",
	})
}

func updateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update An Account",
	})
}

func deleteAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete An Account By ID : " + id,
	})
}
