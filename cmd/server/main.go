package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/chrisarmitage/poc-gotth-skeleton/internal/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/", handleIndex())
	mux.Get("/records", handleListRecords())

	slog.Info("starting application", "addr", "localhost:3000")

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		slog.Error("fatal error", "err", err)
	}
}

func handleIndex() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index()).ServeHTTP(w, r)
	}
}

func handleListRecords() func(w http.ResponseWriter, r *http.Request) {
	l := []string{
		"Item Alpha",
		"Item Bravo",
		"Item Charlie",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.List(l)).ServeHTTP(w, r)
	}
}
