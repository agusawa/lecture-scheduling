package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = New(TEST_ENV_PATH)

func TestSqlite_NewSqliteDatabase(t *testing.T) {
	t.Run("Should create database file", func(t *testing.T) {
		pathname := fmt.Sprintf("../%s", config.Get("SQL_FILENAME"))
		os.Remove(pathname)

		assert.NoFileExists(t, pathname)

		NewSqliteDatabase(config)
		assert.FileExists(t, pathname)
	})
}
