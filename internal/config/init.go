package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func loadEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

type Config struct {
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	DB_URL             string
	ACCESS_KEY         string
	ACCESS_TIME        time.Duration
	REFRESH_KEY        string
	REFRESH_TIME       time.Duration
	APP_VERSION        string
	PORT               string
	GIN_MODE           string
	LOGGER_FOLDER_PATH string
	LOGGER_FILENAME    string
	UPLOAD_PATH        string
}

var ENV Config

func Init() *Config {
	godotenv.Load(".env")

	ENV.PORT = loadEnvVariable("PORT")

	ENV.DB_HOST = loadEnvVariable("DB_HOST")
	ENV.DB_PORT = loadEnvVariable("DB_PORT")
	ENV.DB_USER = loadEnvVariable("DB_USER")
	ENV.DB_PASSWORD = loadEnvVariable("DB_PASSWORD")
	ENV.DB_NAME = loadEnvVariable("DB_NAME")

	ENV.GIN_MODE = loadEnvVariable("GIN_MODE")

	ENV.LOGGER_FOLDER_PATH = loadEnvVariable("LOGGER_FOLDER_PATH")
	ENV.LOGGER_FILENAME = loadEnvVariable("LOGGER_FILENAME")

	ENV.ACCESS_KEY = loadEnvVariable("ACCESS_KEY")
	AT, _ := time.ParseDuration(loadEnvVariable(("ACCESS_TIME")))
	ENV.ACCESS_TIME = AT
	ENV.REFRESH_KEY = loadEnvVariable("REFRESH_KEY")
	RT, _ := time.ParseDuration(loadEnvVariable(("REFRESH_TIME")))
	ENV.REFRESH_TIME = RT

	ENV.APP_VERSION = loadEnvVariable("APP_VERSION")
	ENV.UPLOAD_PATH = loadEnvVariable("UPLOAD_PATH")
	return &ENV
}
