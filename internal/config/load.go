package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    int
	DBUrl   string
	Storage string
	Secret  string
	Expires time.Duration
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

	expires, err := time.ParseDuration(os.Getenv("EXPIRES"))
	if err != nil {
		log.Fatal("Invalid token expiration field")
	}

	cfg := Config{
		Port:    port,
		DBUrl:   os.Getenv("DB_URL"),
		Storage: os.Getenv("STORAGE"),
		Secret:  os.Getenv("SECRET"),
		Expires: expires,
		Env:     os.Getenv("ENV"),
	}

	return &cfg
}
