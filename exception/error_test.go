package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException_PanicIfNeeded(t *testing.T) {
	t.Run("Should panic", func(t *testing.T) {
		err := errors.New("Error message")
		assert.Panics(t, func() { PanicIfNeeded(err) })
	})

	t.Run("Should not panic", func(t *testing.T) {
		var err error = nil
		assert.NotPanics(t, func() { PanicIfNeeded(err) })
	})
}
