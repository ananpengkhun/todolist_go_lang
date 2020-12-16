package controller

import (
	"encoding/json"
	"net/http"

	"github.com/menza01/views"
)

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Page not found",
		}
		json.NewEncoder(w).Encode(data)

	}
}
