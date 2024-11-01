package server

import (
	"net/http"
	"strconv"

	"github.com/Noiidor/go-service-template/internal/domain"
)

func (s *Server) handleWizardsCreate() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wizard, err := decodeRequestBody[domain.Wizard](r)
			if err != nil {
				responseError(w, err, http.StatusBadRequest)
				return
			}

			err = s.WizardService.Create(r.Context(), &wizard)
			if err != nil {
				responseError(w, err, http.StatusInternalServerError)
				return
			}

			encodeResponse(w, http.StatusOK, wizard)
		})
}

func (s *Server) handleWizardsGet() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
			if err != nil {
				responseError(w, err, http.StatusBadRequest)
				return
			}

			wizard, err := s.WizardService.GetByID(r.Context(), uint32(id))
			if err != nil {
				responseError(w, err, http.StatusInternalServerError)
				return
			}

			encodeResponse(w, http.StatusOK, wizard)
		})
}

func (s *Server) handleWizardsList() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wizards, err := s.WizardService.GetAll(r.Context())
			if err != nil {
				responseError(w, err, http.StatusInternalServerError)
				return
			}

			encodeResponse(w, http.StatusOK, wizards)
		})
}

func (s *Server) handleWizardsUpdate() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
			if err != nil {
				responseError(w, err, http.StatusBadRequest)
				return
			}

			update, err := decodeRequestBody[domain.UpdateWizard](r)
			if err != nil {
				responseError(w, err, http.StatusBadRequest)
				return
			}

			wizard, err := s.WizardService.Update(r.Context(), uint32(id), &update)
			if err != nil {
				responseError(w, err, http.StatusInternalServerError)
				return
			}

			encodeResponse(w, http.StatusOK, wizard)
		})
}

func (s *Server) handleWizardsDelete() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
			if err != nil {
				responseError(w, err, http.StatusBadRequest)
				return
			}

			err = s.WizardService.Delete(r.Context(), uint32(id))
			if err != nil {
				responseError(w, err, http.StatusInternalServerError)
				return
			}

			encodeResponse[any](w, http.StatusOK, nil)
		})
}
