package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Payload struct {
	URL string `json:"url" validate:"required"`
}

func (app *application) handleShorten(w http.ResponseWriter, r *http.Request) {
	var payload Payload

	if err := readJSON(r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := validateURL(payload.URL); err != nil {
		app.badRequest(w, r, err)
		return
	}

	shortCode := app.generateRandShortCode()

	entry := &ShortURL{
		URL:       payload.URL,
		ShortCode: shortCode,
	}

	if err := app.storage.URLS.Create(r.Context(), entry); err != nil {
		app.internalError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, entry); err != nil {
		app.internalError(w, r, err)
	}

}

func (app *application) handleGet(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")
	if shortCode == "" {
		app.badRequest(w, r, errors.New("no short code provided"))
		return
	}

	log.Printf("[%v] request", shortCode)
	entry, err := app.storage.URLS.GetByShortCode(r.Context(), shortCode)
	if err != nil {
		switch err {
		case ErrNotFound:
			app.notFound(w, r, err)
		default:
			app.internalError(w, r, err)
		}
		return
	}

	// Update short url stats
	entry.UpdatedAt = time.Now().Format(time.RFC3339)
	entry.AccessCount++

	if err := app.storage.URLS.UpdateStats(r.Context(), entry); err != nil {
		switch err {
		case ErrNotFound:
			app.notFound(w, r, err)
		default:
			app.internalError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, entry); err != nil {
		app.internalError(w, r, err)
	}
}

func (app *application) handleUpdate(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	entry, err := app.storage.URLS.GetByShortCode(r.Context(), shortCode)
	if err != nil {
		switch err {
		case ErrNotFound:
			app.notFound(w, r, err)
		default:
			app.internalError(w, r, err)
		}
		return
	}

	var payload Payload

	if err := readJSON(r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := validateURL(payload.URL); err != nil {
		app.badRequest(w, r, err)
		return
	}

	// Update entry
	entry.URL = payload.URL
	entry.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := app.storage.URLS.Update(r.Context(), entry); err != nil {
		switch err {
		case ErrNotFound:
			app.notFound(w, r, err)
		default:
			app.internalError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, entry); err != nil {
		app.internalError(w, r, err)
	}
}

func (app *application) handleDelete(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	if err := app.storage.URLS.Delete(r.Context(), shortCode); err != nil {
		switch err {
		case ErrNotFound:
			app.notFound(w, r, err)
		default:
			app.internalError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, "short url deleted"); err != nil {
		app.internalError(w, r, err)
	}

}

func (app *application) handleStats(w http.ResponseWriter, r *http.Request) {
	// TODO stats implementation

	if err := app.jsonResponse(w, http.StatusNotFound, "no data"); err != nil {
		app.internalError(w, r, err)
	}
}
