package crud

import (
	"testing"
)

func TestIntGenerator(t *testing.T) {
	var g IDGenerator[int] = &IntGenerator{}

	next, _ := g.Next()

	if next != 1 {
		t.Fatal("next should be 1")
	}

	next, _ = g.Next()

	if next != 2 {
		t.Fatal("next should be 2")
	}
}

func TestStringGenerator(t *testing.T) {
	var s IDGenerator[string] = &StringGenerator{}

	next1, _ := s.Next()
	next2, _ := s.Next()

	if next1 == next2 {
		t.Fatal("next1 and next2 values should be different")
	}
}
