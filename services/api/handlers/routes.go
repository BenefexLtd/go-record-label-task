package handlers

import (
	"net/http"

	"cloud.google.com/go/pubsub"
)

type Application struct {
	PubsubClient *pubsub.Client
}

func (a *Application) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /releases", a.handlePOSTReleases)

	return mux
}
