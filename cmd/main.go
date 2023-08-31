package main

import (
	"avitoTech/config"
	"avitoTech/internal/handlers"
	"avitoTech/internal/infrastructure/server"
	"avitoTech/internal/services"
	"avitoTech/internal/store/sql"
	"avitoTech/internal/utils"
	"avitoTech/tools/db"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Errorf(".env load error: %v\n", err.Error())
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Errorf("config not found: %v", err.Error())
	}

	sqlx, err := db.NewDB(&cfg.DB)
	if err != nil {
		slog.Errorf("db client for sqlx: %v\n", err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	userSegmentStore := sql.New(sqlx)
	userSegmentService := services.NewUserSegmentService(userSegmentStore.Segment())
	userSegmentConteiner := services.New(*userSegmentService)

	srv := createServer(ctx, cancel, cfg, userSegmentConteiner)

	go utils.ManageDBUserSegment(ctx, sqlx)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		c := <-stopChan
		slog.Infof("Caught signal, %v\n", c.String())
		cancel()
	}()

	srv.Start()

	defer func() {
		err = srv.Stop()
		if err != nil {
			slog.Errorf("an error occurred while shutdown server: %v\n", err)
		}
	}()
	<-ctx.Done()

}

func createServer(ctx context.Context,
	cancel context.CancelFunc,

	cfg *config.Config, userSegmentService services.Services) *handlers.Server {

	srv := server.New(ctx, cancel, &cfg.Server, userSegmentService)

	server := handlers.NewServer(ctx, srv, *cfg)
	return server
}
