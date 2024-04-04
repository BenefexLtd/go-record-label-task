package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errBody struct {
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, err error, status int) {
	resp := errBody{Message: err.Error()}

	j, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("cannot encode response", "error", err)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(j); err != nil {
		slog.Error("cannot write response", "error", err)
	}
}
