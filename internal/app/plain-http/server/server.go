package server

import (
	"log/slog"
	"net/http"
)

func NewServer(
	logger *slog.Logger,
	config interface{},
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, config)

	var handler http.Handler = mux

	return handler
}
