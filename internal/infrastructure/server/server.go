package server

import (
	"avitoTech/internal/infrastructure/server/config"
	"avitoTech/internal/infrastructure/server/middleware"
	"avitoTech/internal/services"
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/pkg/errors"
)

type Server struct {
	Ctx    context.Context
	Server *http.Server
	Engine *gin.Engine
	Cancel context.CancelFunc
	Cfg    *config.SrvConfig
}

func New(ctx context.Context, cancel context.CancelFunc, cfg *config.SrvConfig, userSegmentService services.Services) *Server {
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestMiddleware(ctx, userSegmentService))
	srv := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler:      engine,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
	return &Server{
		Ctx:    ctx,
		Cancel: cancel,
		Cfg:    cfg,
		Server: srv,
		Engine: engine,
	}
}

func (s *Server) AddRoutes(routes ...*Route) {
	for _, route := range routes {
		s.Engine.Handle(route.Method, route.Path, route.HandleFuncs...)
	}
}

func (s *Server) Start() {
	slog.Info("Starting HTTP server...")
	//fmt.Println("Starting HTTP server...")
	go func() {
		err := s.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Errorf("listen filed %v", err.Error())

		}
	}()
}

func (s *Server) Stop() error {
	s.Cancel()
	var ctx context.Context
	var cancel context.CancelFunc
	if s.Cfg.ShutdownTimeout != 0 {
		ctx, cancel = context.WithTimeout(s.Ctx, s.Cfg.ShutdownTimeout)

	} else {
		ctx, cancel = context.WithCancel(s.Ctx)
	}
	defer cancel()
	err := s.Server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "shutdown failed")
	}
	return nil
}
