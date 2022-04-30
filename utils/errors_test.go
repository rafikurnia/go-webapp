package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A test to check if the error message matches with what is expected
func TestErrorMessages(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		err := ErrorNotFound
		expect := "not found"

		assert.EqualErrorf(t, err, expect, "Expect %v, got %v", expect, err)
	})

	t.Run("duplicate entry", func(t *testing.T) {
		err := ErrorDuplicateEntry
		expect := "duplicate entry"

		assert.EqualErrorf(t, err, expect, "Expect %v, got %v", expect, err)
	})
}
