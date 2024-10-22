package server

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	config interface{},
) {
	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error("Echo endpoint", "err", err)
			http.Error(w, fmt.Sprintf("err while reading body: %s", err), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

}
