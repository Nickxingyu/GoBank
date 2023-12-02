package router

import (
	"net/http"

	"github.com/Nickxingyu/GoBank/internal/config"
	"github.com/Nickxingyu/GoBank/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	user := User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), config.Config.Bcrypt.Cost)
	if err != nil {
		ctx.Error(err)
		return
	}

	var userModel *model.User
	userModel, err = model.InsertUser(model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Hash:      string(hash),
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	user.ID = userModel.ID

	ctx.JSON(http.StatusOK, user)
}
