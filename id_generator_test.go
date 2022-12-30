package gocrud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntGenerator(t *testing.T) {
	g := NewIntGenerator()

	next, _ := g.Next()

	assert.Equal(t, 1, next)

	next, _ = g.Next()

	assert.Equal(t, 2, next)
}

func TestStringGenerator(t *testing.T) {
	g := NewStringGenerator()

	next1, _ := g.Next()
	next2, _ := g.Next()

	assert.NotEqual(t, next1, next2)
}
