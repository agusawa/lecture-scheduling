package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFlags(t *testing.T) {
	envsTest := []struct {
		name         string
		defaultValue interface{}
	}{
		{"TODAY", false},
		{"MANAGE", false},
	}

	t.Run("Before InitFlags()", func(t *testing.T) {
		for _, envTest := range envsTest {
			assert.Empty(t, os.Getenv(envTest.name))
		}
	})

	t.Run("After InitFlags()", func(t *testing.T) {
		InitFlags()

		for _, envTest := range envsTest {
			switch reflect.TypeOf(envTest.defaultValue).Kind() {
			case reflect.Bool:
				assert.Equal(t, envTest.defaultValue, os.Getenv(envTest.name) == "true")
			}
		}
	})
}
