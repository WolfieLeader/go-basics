package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test ./...
func TestRequestLineParse(t *testing.T) {
	assert.Equal(t, "TheTestagen", "TheTestagen")
}
