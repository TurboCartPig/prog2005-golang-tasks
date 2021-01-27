package fibandfact

import "testing"

func TestFib(t *testing.T) {
	var got [10]int
	for i := 0; i < 10; i++ {
		got[i] = Fib(i + 1)
	}

	expected := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, e := range expected {
		if e != got[i] {
			t.Errorf("Expected %d, got %d", e, got[i])
		}
	}
}

func TestFact(t *testing.T) {
	var got [10]int
	for i := 0; i < 10; i++ {
		got[i] = Fact(i + 1)
	}

	expected := [10]int{1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	for i, e := range expected {
		if e != got[i] {
			t.Errorf("Expected %d, got %d", e, got[i])
		}
	}
}
