package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine) {
	engine.GET("/", greet)
	engine.POST("login", login)
	engine.POST("signUp", signUp)

	UserRouter{engine}.load()
	AccountRouter{engine}.load()
}

func greet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to GoBank.",
	})
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

func signUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sign Up",
	})
}
