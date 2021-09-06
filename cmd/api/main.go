package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("run failed: %s", err.Error())
	}
}

func run() error {
	errC := make(chan error)
	sigC := make(chan os.Signal)
	relC := make(chan Release)

	signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := runServer(":8080", newHandler(relC)); err != nil {
			errC <- err
			return
		}
	}()

	for {
		select {
		case rel := <-relC:
			log.Printf(
				"New release! %s: %s by %s - we should probably let some people know about it...",
				rel.ReleaseDate.Format("02 Jan 2006"),
				rel.Title,
				rel.Artist,
			)
		case err := <-errC:
			return fmt.Errorf("error received: %w", err)
		case <-sigC:
			log.Println("cleaning up...")
			return nil
		}
	}
}

func runServer(addr string, hnd http.Handler) error {
	log.Printf("listening on %s...", addr)

	if err := http.ListenAndServe(addr, hnd); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server failed: %w", err)
	}

	return nil
}

func newHandler(relC chan<- Release) http.Handler {
	hnd := http.NewServeMux()
	hnd.HandleFunc("/releases", releasesHandler(relC))
	return hnd
}

func releasesHandler(relC chan<- Release) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var inp []Release

		if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
			writeError(w, fmt.Errorf("cannot decode request: %w", err), http.StatusBadRequest)
			return
		}

		for _, rel := range inp {
			relC <- rel
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func writeError(w http.ResponseWriter, err error, status int) {
	type errBody struct {
		Message string `json:"message"`
	}

	resp := errBody{Message: err.Error()}
	b := &bytes.Buffer{}

	if err := json.NewEncoder(b).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("cannot encode response: %s", err.Error())
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(b.Bytes()); err != nil {
		log.Printf("cannot write response: %s", err.Error())
	}
}

type Release struct {
	Artist       string                `json:"artist"`
	Title        string                `json:"title"`
	Genre        string                `json:"genre"`
	ReleaseDate  time.Time             `json:"releaseDate"`
	Distribution []ReleaseDistribution `json:"distribution"`
}

type ReleaseDistribution struct {
	Type string `json:"type"`
	Qty  int64  `json:"qty"`
}
