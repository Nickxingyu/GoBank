package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	rg := engine.Group("/user")
	{
		rg.GET("/:id", getUserById)
		rg.POST("/", createUser)
		rg.PUT("/", updateUser)
		rg.DELETE("/:id", deleteUserById)
	}
}

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get A User By ID : " + id,
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
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete A User By ID : " + id,
	})
}
