package main

import "database/sql"

func NewPostgresConn(addr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
