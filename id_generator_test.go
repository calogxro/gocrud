package gocrud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDGenerator(t *testing.T) {
	g := IDGenerator{}

	next, _ := g.Next()

	assert.Equal(t, 1, next)

	next, _ = g.Next()

	assert.Equal(t, 2, next)
}
