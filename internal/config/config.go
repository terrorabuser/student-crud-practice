package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() *Config {

	readTimeout := mustGetEnvAsInt("READ_TIMEOUT")
	writeTimeout := mustGetEnvAsInt("WRITE_TIMEOUT")

	return &Config{
		Server: ServerConfig{
			Port:         mustGetEnv("SERVER_PORT"),
			Host:         mustGetEnv("SERVER_HOST"),
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		},
		Database: DatabaseConfig{
			Driver:   mustGetEnv("DB_DRIVER"),
			Host:     mustGetEnv("DB_HOST"),
			Port:     mustGetEnv("DB_PORT"),
			User:     mustGetEnv("DB_USER"),
			Password: mustGetEnv("DB_PASSWORD"),
			Name:     mustGetEnv("DB_NAME"),
		},
	}
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return val
}

func mustGetEnvAsInt(key string) int {
	valStr := mustGetEnv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatalf("Environment variable %s must be an integer", key)
	}
	return val
}
