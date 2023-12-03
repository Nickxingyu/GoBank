package router

import (
	"net/http"
	"strconv"

	"github.com/Nickxingyu/GoBank/internal/config"
	"github.com/Nickxingyu/GoBank/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserOutput struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func init() {
	rg := engine.Group("/user")
	{
		rg.GET("/:id", getUserById)
		rg.PUT("/", updateUser)
		rg.DELETE("/:id", deleteUserById)
	}
}

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	userModel, err := model.FindUserByID(uint(userID))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, UserOutput{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	})
}

func updateUser(ctx *gin.Context) {
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

	userModel, err := model.SaveUserByID(user.ID, model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Hash:      string(hash),
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, UserOutput{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	})

}

func deleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	userModel, err := model.DeleteUserByID(uint(userID))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, UserOutput{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	})
}
