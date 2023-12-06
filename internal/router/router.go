package router

import (
	"log"
	"net/http"

	"github.com/Nickxingyu/GoBank/internal/config"
	"github.com/Nickxingyu/GoBank/internal/model"
	"github.com/Nickxingyu/GoBank/internal/service/jwt"
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
	user := User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		return
	}

	userModel, err := model.FindUserByEmail(user.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Hash), []byte(user.Password)); err != nil {
		log.Println("Invalid Password")
		ctx.Error(err)
		return
	}

	tokenString, err := jwt.GenerateTokenString(userModel.ID)
	if err != nil {
		log.Println("JWT Process Fail")
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   tokenString,
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

	userModel, err := model.InsertUser(model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Hash:      string(hash),
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, User{
		ID:        userModel.ID,
		FirstName: user.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	})
}
