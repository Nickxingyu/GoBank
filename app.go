package main

import (
	"github.com/Nickxingyu/GoBank/internal/app"
	"github.com/Nickxingyu/GoBank/internal/model"
)

func main() {
	model.ConnectDatabase()
	server := app.NewApiServer(":3000")
	server.Run()
}
