package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	Host       string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBName:     getEnv("DB_NAME", "mydb"),
		Port:       getEnv("PORT", "8080"),
		Host:       getEnv("HOST", "localhost"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
