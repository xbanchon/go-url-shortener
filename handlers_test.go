package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetShortURL(t *testing.T) {
	app := newTestApp(t, config{})
	mux := app.mount()

	t.Run("should not allow empty short code", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should not allow short code with invalid characters", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/:dl,ffmk", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should not accept short code of invalid length", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten/utFl", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})
}

func TestCreateURL(t *testing.T) {
	app := newTestApp(t, config{})
	mux := app.mount()

	t.Run("should not allow shorten if long URL is not provided", func(t *testing.T) {
		payload := Payload{
			URL: "",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/shorten/", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should not allow shorten if long URL is malformed", func(t *testing.T) {
		payload := Payload{
			URL: "randomstring",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/shorten/", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})
}
