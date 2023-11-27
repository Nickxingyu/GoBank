package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine) {
	engine.GET("/", greet)
	AccountRouter{engine}.Load()
}

func greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GoBank.",
	})
}
