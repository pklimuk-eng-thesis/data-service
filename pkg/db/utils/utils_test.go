package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tm, err := ParseTime("2023-01-01T00:00:00Z")
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), tm)
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := ParseTime("Wrong data")
		assert.Error(t, err)
	})
}
