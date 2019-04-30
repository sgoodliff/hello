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
// check that GoodBye() returns expected
func TestGoodBye(t *testing.T) {
	expected := "GoodBye!"
	actual := Hello()
	if actual != expected {
		t.Error("Test failed")
	}
}