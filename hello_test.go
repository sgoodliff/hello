package main

import "testing"

// check that hello() returns expected
func TestHello(t *testing.T) {
	expected := "Hello Go!"
	actual := Hello()
	if actual != expected {
		t.Error("Test failed")
	}
}
