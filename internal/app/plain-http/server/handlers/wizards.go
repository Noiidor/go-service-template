package handlers

import "net/http"

func HandleWizardsCreate() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
		})
}
