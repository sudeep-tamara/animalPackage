package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(4, 5)
	if total != 9 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 9)
	}
}
