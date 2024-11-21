package main

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrInternal   = errors.New("internal error")
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
)

func (app *application) internalError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%v %v error: %v", r.Method, r.RequestURI, err.Error())
	writeJSONError(w, http.StatusInternalServerError, ErrInternal.Error())
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%v %v error: %v", r.Method, r.RequestURI, err.Error())
	writeJSONError(w, http.StatusBadRequest, ErrBadRequest.Error())
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%v %v error: %v", r.Method, r.RequestURI, err.Error())
	writeJSONError(w, http.StatusNotFound, ErrNotFound.Error())
}
