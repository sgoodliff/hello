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

func TestWhatsup(t *testing.T) {
	expected := "Whatsup!"
	actual := Whatsup()
	if actual != expected {
		t.Error("Test failed")
	}
}

// check that Goodbye() returns expected
func TestBye(t *testing.T) {
	expected := "bye bye"
	actual := Bye()
	if actual != expected {
		t.Error("Test failed")
	}
}
