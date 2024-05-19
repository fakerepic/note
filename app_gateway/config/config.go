package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	COUCHDB_USER     string
	COUCHDB_PASSWORD string
	COUCHDB_HOST     string
	COUCHDB_PORT     string
	AISERVICE_HOST   string
	AISERVICE_PORT   string
	ADMIN_EMAIL      string
	ADMIN_PASSWORD   string
	SMTP_HOST        string
	SMTP_PORT        string
	SMTP_USER        string
	SMTP_PASSWORD    string
	APP_PUBLIC_URL   string
}

func Load() Config {
	godotenv.Load()

	return Config{
		COUCHDB_USER:     os.Getenv("COUCHDB_USER"),
		COUCHDB_PASSWORD: os.Getenv("COUCHDB_PASSWORD"),
		COUCHDB_HOST:     os.Getenv("COUCHDB_HOST"),
		COUCHDB_PORT:     os.Getenv("COUCHDB_PORT"),
		AISERVICE_HOST:   os.Getenv("AISERVICE_HOST"),
		AISERVICE_PORT:   os.Getenv("AISERVICE_PORT"),
		ADMIN_EMAIL:      os.Getenv("ADMIN_EMAIL"),
		ADMIN_PASSWORD:   os.Getenv("ADMIN_PASSWORD"),
		SMTP_HOST:        os.Getenv("SMTP_HOST"),
		SMTP_PORT:        os.Getenv("SMTP_PORT"),
		SMTP_USER:        os.Getenv("SMTP_USER"),
		SMTP_PASSWORD:    os.Getenv("SMTP_PASSWORD"),
		APP_PUBLIC_URL:   os.Getenv("APP_PUBLIC_URL"),
	}
}

func (c Config) CouchAdminUrl() string {
	return fmt.Sprintf("http://%s:%s@%s:%s", c.COUCHDB_USER, c.COUCHDB_PASSWORD, c.COUCHDB_HOST, c.COUCHDB_PORT)
}

func (c Config) CouchUrl() string {
	return fmt.Sprintf("http://%s:%s", c.COUCHDB_HOST, c.COUCHDB_PORT)
}

func (c Config) AIServiceUrl() string {
	return fmt.Sprintf("http://%s:%s", c.AISERVICE_HOST, c.AISERVICE_PORT)
}
