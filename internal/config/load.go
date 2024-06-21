package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    int
	DBUrl   string
	Storage string
	Env     string
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid port")
	}

	cfg := Config{
		Port:    port,
		DBUrl:   os.Getenv("DB_URL"),
		Storage: os.Getenv("STORAGE"),
		Env:     os.Getenv("ENV"),
	}

	return &cfg
}
