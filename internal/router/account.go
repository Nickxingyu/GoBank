package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Nickxingyu/GoBank/internal/model"
)

type Account struct {
	ID       uint    `json:"id"`
	UserID   uint    `json:"user_id"`
	CoinType string  `json:"coin_type"`
	Balance  float64 `json:"balance"`
}

func init() {
	rg := engine.Group("/account")
	{
		rg.GET("/:id", getAccountById)
		rg.POST("/", createAccount)
		rg.PUT("/", updateAccount)
		rg.DELETE("/:id", deleteAccountById)
	}
}

func getAccountById(ctx *gin.Context) {
	id := ctx.Param("id")

	account_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	accountModel, err := model.FindAccountByID(uint(account_id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &Account{
		ID:       accountModel.ID,
		UserID:   accountModel.UserID,
		CoinType: accountModel.CoinType,
		Balance:  accountModel.Balance,
	})
}

func createAccount(ctx *gin.Context) {
	account := Account{}
	err := ctx.BindJSON(&account)
	if err != nil {
		ctx.Error(err)
		return
	}

	_, err = model.InsertAccount(model.Account{
		UserID:   account.UserID,
		CoinType: account.CoinType,
		Balance:  account.Balance,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func updateAccount(ctx *gin.Context) {
	account := Account{}
	err := ctx.BindJSON(&account)
	if err != nil {
		ctx.Error(err)
		return
	}

	_, err = model.SaveAccountByID(account.ID, model.Account{
		UserID:   account.UserID,
		CoinType: account.CoinType,
		Balance:  account.Balance,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func deleteAccountById(ctx *gin.Context) {
	id := ctx.Param("id")

	account_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	accountModel, err := model.DeleteAccountByID(uint(account_id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &Account{
		ID:       accountModel.ID,
		UserID:   accountModel.UserID,
		CoinType: accountModel.CoinType,
		Balance:  accountModel.Balance,
	})
}
