package main

import "testing"

func TestParadise(t *testing.T) {
	got := paradise("Munich")
	want := "My idea of paradise is Munich"
	if got != want {
		t.Errorf("Error in paradise function, got: %s, want: %s", got, want)
	}

}
