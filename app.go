package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	cfg     config
	storage *sql.DB
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/shorten", func(r chi.Router) {
		r.Post("/", app.handleShorten)

		r.Route("/{shortCode}", func(r chi.Router) {
			r.Get("/", app.handleGet)
			r.Put("/", app.handleUpdate)
			r.Delete("/", app.handleDelete)
			r.Get("/stats", app.handleStats)
		})
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.cfg.serverCfg.addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("server listening in port %v...", app.cfg.serverCfg.addr)
	return srv.ListenAndServe()
}
