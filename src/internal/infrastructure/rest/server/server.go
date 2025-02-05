package server

import (
	"fmt"
	"receipt-api/src/internal/infrastructure/rest/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
}

func NewServer(router *router.Router) *Server {
	engine := gin.New()

	return &Server{
		engine: engine,
		router: router,
	}
}

func (s *Server) SetupRoutes() {
	s.engine.Use(gin.Recovery())

	s.router.SetupRoutes(s.engine)
}

func (s *Server) Start(port string) error {
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
