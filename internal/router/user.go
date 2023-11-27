package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	*gin.Engine
}

func (r UserRouter) load() {
	rg := r.Group("/user")

	rg.GET("/:id", getUserById)
	rg.POST("/", createUser)
	rg.PUT("/", updateUser)
	rg.DELETE("/:id", deleteUserById)
}

func getUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get A User By Id",
	})
}

func createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create A User",
	})
}

func updateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update A User",
	})
}

func deleteUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete A User By Id",
	})
}