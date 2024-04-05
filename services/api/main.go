package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"

	"github.com/BenefexLtd/go-record-label-task/services/api/handlers"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Starting API service...")

	pc, err := pubsub.NewClient(ctx, "benefex-test")
	if err != nil {
		slog.Error("creating new pubsub client failed", "error", err)
		os.Exit(1)
	}

	if err := createPubsubTopics(ctx, pc, []string{"releases"}); err != nil {
		slog.Error("creating new pubsub topics failed", "error", err)
		os.Exit(1)
	}

	app := handlers.Application{
		PubsubClient: pc,
	}
	handler := app.InitRoutes()
	slog.Info("Listening for HTTP requests on :8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("running http server failed", "error", err)
		os.Exit(1)
	}
}

func createPubsubTopics(ctx context.Context, client *pubsub.Client, topics []string) error {
	for _, topic := range topics {
		exists, err := client.Topic(topic).Exists(ctx)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		if _, err := client.CreateTopic(ctx, topic); err != nil {
			return err
		}
	}
	return nil
}
