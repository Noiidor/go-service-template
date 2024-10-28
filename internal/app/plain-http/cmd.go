package plainhttp

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Noiidor/go-service-template/internal/app/plain-http/server"
	"github.com/Noiidor/go-service-template/internal/config"
	"github.com/Noiidor/go-service-template/internal/db/postgres"
	postgresrepos "github.com/Noiidor/go-service-template/internal/repos/postgres"
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

	srv := server.NewServer(logger, cfg)

	httpServer := http.Server{
		Handler: srv,
		Addr:    net.JoinHostPort("", "5050"),
	}

	//// Graceful shutdown

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(httpServer.ListenAndServe)
	eg.Go(func() error {
		<-egCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout*time.Second)
		defer cancel()

		return httpServer.Shutdown(shutdownCtx)
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
