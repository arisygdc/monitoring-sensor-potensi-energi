package server

import (
	"monitoring-potensi-energi/config"
	"monitoring-potensi-energi/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env    config.Environment
	engine *gin.Engine
	ctr    controller.Controller
}

func New(env config.Environment, controller controller.Controller) (server Server) {
	engine := gin.Default()
	gin.SetMode(server.env.ServerEnv)
	server = Server{
		env:    env,
		engine: engine,
		ctr:    controller,
	}
	return
}

func (s *Server) Run() {
	s.engine.Run(s.env.ServerAddress)
}
