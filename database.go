package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"
)

var (
	QueryTimeout = 5 * time.Second
)

type ShortURL struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	ShortCode   string `json:"shortCode"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	AccessCount int    `json:"accessCount"`
}

func NewStorage(addr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) CreateURL(ctx context.Context, entry *ShortURL) error {
	query := `
		INSERT INTO urls (url, short_code)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at, access_count
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	err := app.storage.QueryRowContext(
		ctx,
		query,
		entry.URL,
		entry.ShortCode,
	).Scan(
		&entry.ID,
		&entry.CreatedAt,
		&entry.UpdatedAt,
		&entry.AccessCount,
	)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) GetURLByShortCode(ctx context.Context, code string) (*ShortURL, error) {
	query := `
		SELECT id, url, short_code, created_at, updated_at, access_count
		FROM urls
		WHERE short_code = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	entry := &ShortURL{}

	err := app.storage.QueryRowContext(
		ctx,
		query,
		code,
	).Scan(
		&entry.ID,
		&entry.URL,
		&entry.ShortCode,
		&entry.CreatedAt,
		&entry.UpdatedAt,
		&entry.AccessCount,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return entry, nil
}

func (app *application) UpdateURL(ctx context.Context, entry *ShortURL) error {
	query := `
		UPDATE urls
		SET url = $1, updated_at = $2
		WHERE id = $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	res, err := app.storage.ExecContext(
		ctx,
		query,
		entry.URL,
		entry.UpdatedAt,
		entry.ID,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (app *application) UpdateStats(ctx context.Context, entry *ShortURL) error {
	query := `
		UPDATE urls
		SET updated_at = $1, access_count = $2
		WHERE id = $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	res, err := app.storage.ExecContext(
		ctx,
		query,
		entry.UpdatedAt,
		entry.AccessCount,
		entry.ID,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (app *application) DeleteURL(ctx context.Context, code string) error {
	query := `
			DELETE FROM urls
			WHERE short_code = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	res, err := app.storage.ExecContext(
		ctx,
		query,
		code,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
