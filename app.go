package main

import (
	"github.com/Nickxingyu/GoBank/internal/app"
)

func main() {
	server := app.NewApiServer(":3000")
	server.Run()
}
