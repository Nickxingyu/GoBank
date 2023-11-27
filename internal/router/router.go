package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine) {
	engine.GET("/", greet)

	accountRouters := engine.Group("/account")
	{
		accountRouters.GET("/:id", getAccountById)
		accountRouters.POST("/", createAccount)
		accountRouters.PUT("/", updateAccount)
		accountRouters.DELETE("/:id", deleteAccountById)
	}
}

func greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GoBank.",
	})
}
