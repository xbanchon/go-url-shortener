package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApp(t *testing.T, cfg config) *application {
	t.Helper()
	mockStorage := newMockStorage()

	return &application{
		cfg:     cfg,
		storage: mockStorage,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
