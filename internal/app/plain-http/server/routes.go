package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.log.Error("Echo endpoint", "err", err)
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

	mux.Handle("POST /v1/wizards", s.handleWizardsCreate())
	mux.Handle("PATCH /v1/wizards/{id}", s.handleWizardsUpdate())
	mux.Handle("GET /v1/wizards", s.handleWizardsList())
	mux.Handle("GET /v1/wizards/{id}", s.handleWizardsGet())
	mux.Handle("DELETE /v1/wizards/{id}", s.handleWizardsDelete())

}
