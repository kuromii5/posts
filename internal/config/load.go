package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	Storage  string
	Env      string
	Postgres PostgresConfig
	Redis    RedisConfig
}

type PostgresConfig struct {
	URL string
}

type RedisConfig struct {
	URL string
}

func MustLoad() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 8000)

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic("error unmarshalling config")
	}

	portStr := viper.GetString("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("invalid port value")
	}
	config.Port = port

	config.Redis.URL = os.Getenv("REDIS_URL")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.URL = fmt.Sprintf("postgres://postgres:%s@postgres:5432/reddit_clone?sslmode=disable", pgPassword)

	return &config
}
