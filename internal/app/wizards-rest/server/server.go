package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Noiidor/go-service-template/internal/config"
	"github.com/Noiidor/go-service-template/internal/service"
)

type Server struct {
	ctx context.Context

	log *slog.Logger
	cfg *config.Config

	server *http.Server

	WizardService service.WizardsService
}

func NewServer(
	ctx context.Context,
	logger *slog.Logger,
	cfg *config.Config,
	wizardService service.WizardsService,
) *Server {
	server := &Server{
		ctx:           ctx,
		log:           logger,
		cfg:           cfg,
		WizardService: wizardService,
	}

	mux := http.NewServeMux()
	server.addRoutes(mux)

	server.server = &http.Server{
		Addr: net.JoinHostPort(
			cfg.GetAppHost(),
			strconv.FormatUint(uint64(cfg.GetAppWizardsRestPort()), 10),
		),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	return server
}

func (s *Server) ListenAndServe() error {
	s.log.Info("HTTP Server started and listening", slog.String("host", s.server.Addr))
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(timeoutCtx context.Context) error {
	s.log.Info("HTTP Server shutting down")
	return s.server.Shutdown(timeoutCtx)
}
