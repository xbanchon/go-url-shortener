package main

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	addr  string
	dbCfg dbConfig
}

type dbConfig struct {
	addr string
}

func loadConfig() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, err
	}

	cfg := config{
		addr: os.Getenv("ADDR"),
		dbCfg: dbConfig{
			addr: os.Getenv("DB_ADDR"),
		},
	}
	return cfg, nil
}
