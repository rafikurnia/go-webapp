package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A test to check that the constants should have different values
func TestDifferentConstants(t *testing.T) {
	t.Run("all constants are different", func(t *testing.T) {
		assert.NotEqualf(t, DBMode, DBModeMock, "Expect %v to be different with %v", DBMode, DBModeMock)
		assert.NotEqualf(t, DBMode, DBModeTest, "Expect %v to be different with %v", DBMode, DBModeTest)
		assert.NotEqualf(t, DBModeMock, DBModeTest, "Expect %v to be different with %v", DBModeMock, DBModeTest)
	})
}
