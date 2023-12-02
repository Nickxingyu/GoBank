package app

import (
	"strconv"

	"github.com/Nickxingyu/GoBank/internal/config"
	"github.com/Nickxingyu/GoBank/internal/router"
)

func Run() {
	port := strconv.Itoa(config.Config.Server.Port)
	router.Run(":" + port)
}
