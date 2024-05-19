package core

import (
	"os"
)

type Config struct {
	DBUri        string
	SMTPUsername string
	SMTPPassword string
	SMTPHost     string
	RedisAddr    string
}

func LoadConfig() Config {
	config := Config{
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		SMTPHost:     os.Getenv("SMTP_HOST"),
		DBUri:        os.Getenv("DB_URI"),
		RedisAddr:    os.Getenv("REDIS_ADDR"),
	}

	return config
}
