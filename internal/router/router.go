package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine = gin.Default()

func init() {
	engine.GET("/", greet)
	engine.POST("/login", login)
	engine.POST("/signUp", signUp)
}

func Run(port string) {
	engine.Run(port)
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
