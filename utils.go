package main

import (
	"errors"
	"math/rand"
	"net/url"
	"time"
)

func (app *application) generateRandShortCode() string {
	charset := app.cfg.shortenerCfg.charset
	n := app.cfg.shortenerCfg.codeLength

	rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func validateURL(inputURL string) error {
	parsedURL, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("invalid scheme")
	}

	return nil
}
