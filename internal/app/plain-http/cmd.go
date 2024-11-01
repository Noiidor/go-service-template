package plainhttp

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Noiidor/go-service-template/internal/app/plain-http/server"
	"github.com/Noiidor/go-service-template/internal/config"
	"github.com/Noiidor/go-service-template/internal/db/postgres"
	postgresrepos "github.com/Noiidor/go-service-template/internal/repos/postgres"
	"github.com/Noiidor/go-service-template/internal/service"
	"golang.org/x/sync/errgroup"
)

const shutdownTimeout = 10

func Run(stdout, stderr io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := slog.New(slog.NewTextHandler(stdout, nil))
	logger.Info("Starting...")

	// Services initialization goes here

	// Loads env variables to config
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	postgres, err := postgres.New(ctx, cfg)
	if err != nil {
		return err
	}

	wizardsRepo := postgresrepos.NewWizardsRepo(postgres)
	_ = wizardsRepo

	wizardService := service.NewWizardsService(logger, wizardsRepo)

	server := server.NewServer(ctx, logger, cfg, wizardService)

	//// Graceful shutdown

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(server.ListenAndServe)
	eg.Go(func() error {
		<-egCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()

		return server.Shutdown(shutdownCtx)
	})

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

		<-exit
		logger.Info("Gracefully shutting down...")
		cancel()
	}()

	logger.Info("HTTP server started!")
	err = eg.Wait()
	logger.Info("Graceful shutdown complete!")
	if err != nil {
		logger.Info("Exit", "reason", err)
	}

	return err
}
