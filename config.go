package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	serverCfg    serverConfig
	dbCfg        dbConfig
	shortenerCfg shortenerConfig
}

type serverConfig struct {
	addr string
}

type dbConfig struct {
	addr string
}

type shortenerConfig struct {
	charset    string
	codeLength int
}

func loadConfig() (config, error) {
	err := godotenv.Load()
	if err != nil {
		return config{}, err
	}

	cfg := config{
		serverCfg: serverConfig{
			addr: os.Getenv("ADDR"),
		},
		dbCfg: dbConfig{
			addr: os.Getenv("DB_ADDR"),
		},
		shortenerCfg: shortenerConfig{
			charset:    os.Getenv("CHARSET"),
			codeLength: getInt("CODE_LENGTH", 6),
		},
	}
	return cfg, nil
}

func getInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return int(intVal)
	}

	return fallback
}
