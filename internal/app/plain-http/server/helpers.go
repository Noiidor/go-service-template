package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encodeResponse[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func decodeRequestBody[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}

func responseError(w http.ResponseWriter, err error, code int) {
	type httpError struct {
		Error string `json:"error"`
	}

	serverError := httpError{
		Error: err.Error(),
	}

	if encodeErr := encodeResponse(w, code, serverError); encodeErr != nil {
		panic(encodeErr)
	}

}
