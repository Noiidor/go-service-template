package server

import (
	"log/slog"
	"net/http"

	"github.com/Noiidor/go-service-template/internal/config"
)

func NewServer(
	logger *slog.Logger,
	cfg *config.Config,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, cfg)

	var handler http.Handler = mux

	return handler
}
