package server

import (
	"context"
	"fmt"
	"log"

	"github.com/badcode256/example_go_hexagonal/internal/infra/server/handler/user"
	"github.com/badcode256/example_go_hexagonal/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr    string
	engine      *gin.Engine
	userService service.UserService
}

func New(ctx context.Context, host string, port uint, userService service.UserService) Server {
	server := Server{
		engine:      gin.Default(),
		httpAddr:    fmt.Sprintf("%s:%d", host, port),
		userService: userService,
	}

	server.routes()
	return server
}

func (s *Server) routes() {
	s.engine.POST("/user/create", user.CreateHandler(s.userService))
}

func (s *Server) Run() error {
	err := s.engine.Run(s.httpAddr)
	if err != nil {
		return err
	}
	log.Println("Server running on", s.httpAddr)

	return nil
}
