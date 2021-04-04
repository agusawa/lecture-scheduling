package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var envs = []struct {
	name         string
	value        string
	defaultValue string
}{
	{"SQL_FILENAME", "database_test.sql", "database.sql"},
}

func resetEnv() {
	for _, env := range envs {
		os.Unsetenv(env.name)
	}
}

func TestConfig_New(t *testing.T) {
	resetEnv()

	t.Run("Get env variable before call the method", func(t *testing.T) {
		for _, env := range envs {
			assert.Empty(t, os.Getenv(env.name))
		}
	})

	resetEnv()

	t.Run("Call with invalid env path", func(t *testing.T) {
		New(".env.wrong")
		for _, env := range envs {
			assert.Equal(t, env.defaultValue, os.Getenv(env.name))
		}
	})

	resetEnv()

	t.Run("Call with correct path", func(t *testing.T) {
		New(TEST_ENV_PATH)
		for _, env := range envs {
			assert.Equal(t, env.value, os.Getenv(env.name))
		}
	})
}

func TestConfig_Get(t *testing.T) {
	config := New(TEST_ENV_PATH)

	t.Run("Get env variable", func(t *testing.T) {
		for _, env := range envs {
			assert.Equal(t, env.value, config.Get(env.name))
		}
	})
}

func TestConfig_DeleteDatabase(t *testing.T) {
	config := New(TEST_ENV_PATH)
	filename := config.Get("SQL_FILENAME")

	t.Run("Delete sql database", func(t *testing.T) {
		file, _ := os.Create(filename)
		file.Close()
		assert.FileExists(t, filename)

		config.DeleteDatabase()
		assert.NoFileExists(t, filename)
	})
}
