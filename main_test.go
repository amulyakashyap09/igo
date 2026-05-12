package igo

import "testing"

func TestGreeting(t *testing.T) {
	name := "Gopher"
	expectedName := "Hello, Gopher!"
	result := "Hello, " + name + "!"

	if result != expectedName {
		t.Errorf("got %q, wanted %q", result, expectedName)
	}
}
