package main

import (
	"net/http"
	"testing"
)

func TestGetShortURL(t *testing.T) {
	app := newTestApp(t, config{})
	mux := app.mount()

	t.Run("Should not allow empty short code", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("Should not allow short code with invalid characters", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/:dl,ffmk", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("Should not accept short code of invalid length", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/utFl", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})
}

// func TestCreateURL(t *testing.T) {}
