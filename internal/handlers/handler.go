package handlers

import (
	"avitoTech/config"
	"avitoTech/internal/infrastructure/server"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*server.Server
	ctx    context.Context
	config config.Config
}

func NewServer(c context.Context, srv *server.Server, cfg config.Config) *Server {
	s := &Server{
		Server: srv,
		ctx:    c,
		config: cfg,
	}
	InjectUserSegmentRoutes(s)
	return s
}

func (s *Server) PingPong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
