package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := add(10, 243)
	if sum != 253 {
		t.Errorf("Sum was incorrect, got %d, want %d", sum, 253)
	}
}

func TestSubtract(t *testing.T) {
	diff := subtract(200, 100)
	if diff != 100 {
		t.Errorf("Diff incorrect, got: %d, expected: %d", diff, 100)
	}
}

func TestDoMath(t *testing.T) {
	result := doMath(10, 5, add)
	if result != 15 {
		t.Errorf("doMath with add was incorrect, got %d, want %d", result, 15)
	}

	result = doMath(10, 5, subtract)
	if result != 5 {
		t.Errorf("doMath with subtract was incorrect, got %d, want %d", result, 5)
	}
}
