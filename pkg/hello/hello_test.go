package hello

import "testing"

// check that hello() returns expected
func TestHello(t *testing.T) {
	expected := "Hello Go!"
	actual := Hello()
	if actual != expected {
		t.Error("Test failed")
	}
}

// check that Goodbye() returns expected
func TestGoodbye(t *testing.T) {
	expected := "Goodbye!"
	actual := Goodbye()
	if actual != expected {
		t.Error("Test failed")
	}
}
func TestGoing(t *testing.T) {
	expected := "Going!"
	actual := Going()
	if actual != expected {
		t.Error("Test failed")
	}
}
