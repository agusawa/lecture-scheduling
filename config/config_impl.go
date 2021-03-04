package config

import (
	"lecture-scheduling/exception"
	"os"

	"github.com/joho/godotenv"
)

func New(envPath string) Config {
	// If dotenv file does not exists
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		loadDefaultEnv()
	} else {
		err := godotenv.Load(envPath)
		exception.PanicIfNeeded(err)
	}

	return &configImpl{}
}

func loadDefaultEnv() {
	os.Setenv("SQL_FILENAME", "database.sql")
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}
