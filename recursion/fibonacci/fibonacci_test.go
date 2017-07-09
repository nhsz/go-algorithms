package fibonacci

import (
	"testing"
)

// Input: 0
// Expected: 0
func TestRecursiveFib1(t *testing.T) {
	if result := RecursiveFib(0); result != 0 {
		t.Errorf("Result was incorrect, got: %v, want: 0.", result)
	}
}

// Input: 1
// Expected: 1
func TestRecursiveFib2(t *testing.T) {
	if result := RecursiveFib(1); result != 1 {
		t.Errorf("Result was incorrect, got: %v, want: 1.", result)
	}
}

// Input: 9
// Expected: 34
func TestRecursiveFib3(t *testing.T) {
	if result := RecursiveFib(9); result != 34 {
		t.Errorf("Result was incorrect, got: %v, want: 34.", result)
	}
}

// Input: 11
// Expected: 89
func TestRecursiveFib4(t *testing.T) {
	if result := RecursiveFib(11); result != 89 {
		t.Errorf("Result was incorrect, got: %v, want: 89.", result)
	}
}

// Input: 0
// Expected: 0
func TestIterativeFib1(t *testing.T) {
	if result := IterativeFib(0); result != 0 {
		t.Errorf("Result was incorrect, got: %v, want: 89.", result)
	}
}

// Input: 1
// Expected: 1
func TestIterativeFib2(t *testing.T) {
	if result := IterativeFib(1); result != 1 {
		t.Errorf("Result was incorrect, got: %v, want: 1.", result)
	}
}

// Input: 9
// Expected: 34
func TestIterativeFib3(t *testing.T) {
	if result := IterativeFib(9); result != 34 {
		t.Errorf("Result was incorrect, got: %v, want: 89.", result)
	}
}

// Input: 11
// Expected: 89
func TestIterativeFib4(t *testing.T) {
	if result := IterativeFib(11); result != 89 {
		t.Errorf("Result was incorrect, got: %v, want: 89.", result)
	}
}
