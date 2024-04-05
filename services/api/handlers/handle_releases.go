package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"cloud.google.com/go/pubsub"

	"github.com/BenefexLtd/go-record-label-task/services/api/domain"
)

func (a *Application) handlePOSTReleases(w http.ResponseWriter, req *http.Request) {
	var releases []domain.Release

	// decode the body
	if err := json.NewDecoder(req.Body).Decode(&releases); err != nil {
		slog.Error("decoding request failed", "error", err)
		wrappedErr := fmt.Errorf("failed to decode request: %w", err)
		writeError(w, wrappedErr, http.StatusBadRequest)
		return
	}

	// bare minimum event publishing for each release
	topic := a.PubsubClient.Topic("releases")
	for _, release := range releases {
		slog.Info("New release! we should probably let some people know about it...",
			"date", release.ReleaseDate.Format("02 Jan 2006"),
			"title", release.Title,
			"artist", release.Artist,
		)

		j, err := json.Marshal(release)
		if err != nil {
			slog.Error("encoding release failed", "error", err)
			wrappedErr := fmt.Errorf("failed to encode release: %w", err)
			writeError(w, wrappedErr, http.StatusInternalServerError)
			return
		}

		res := topic.Publish(req.Context(), &pubsub.Message{Data: j})
		if _, err := res.Get(req.Context()); err != nil {
			slog.Error("publishing event failed", "error", err)
			wrappedErr := fmt.Errorf("failed to publish event: %w", err)
			writeError(w, wrappedErr, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}
