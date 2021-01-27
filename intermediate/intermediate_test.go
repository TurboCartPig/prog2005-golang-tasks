package intermediate

import "testing"

// Task 3

func TestHex(t *testing.T) {
	got := Hex("abcdef")
	expected := "101112131415"
	if got != expected {
		t.Errorf("Expected %s, Got %s", expected, got)
	}
}

func TestHex2(t *testing.T) {
	got := Hex2("abcdef")
	expected := "101112131415"
	if got != expected {
		t.Errorf("Expected %s, Got %s", expected, got)
	}
}
