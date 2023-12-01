package app

import (
	"github.com/Nickxingyu/GoBank/internal/router"
)

var port string

func init() {
	port = ":3000"
}

func Run() {
	router.Run(port)
}
