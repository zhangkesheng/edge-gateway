package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHmacSha256Sign(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		assert.Equal(t, "HCbG3xNE3vzhO+u7qCUL1jS5hsu2n5r2cFhnTrtyDAE=", HmacSha256Sign("testappSecret", "1546084445901"))
	})
}
