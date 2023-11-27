package app

import (
	"github.com/gin-gonic/gin"

	"github.com/Nickxingyu/GoBank/internal/router"
)

type ApiServer struct {
	engine *gin.Engine
	port   string
}

func NewApiServer(port string) *ApiServer {
	engine := gin.Default()

	router.Load(engine)

	return &ApiServer{
		engine: engine,
		port:   port,
	}
}

func (s *ApiServer) Run() {
	engine := s.engine
	port := s.port
	engine.Run(port)
}
