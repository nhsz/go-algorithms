package linearsearch

import (
	"testing"
)

// Input: [], 3
// Expected: -1
func TestLinearSearch1(t *testing.T) {
	list := []int{}
	value := 3

	if position := Search(list, value); position != -1 {
		t.Errorf("Result was incorrect, got: %v, want: -1.", position)
	}
}

// Input: [1], 3
// Expected: -1
func TestLinearSearch2(t *testing.T) {
	list := []int{1}
	value := 3

	if position := Search(list, value); position != -1 {
		t.Errorf("Result was incorrect, got: %v, want: -1.", position)
	}
}

// Input: [1, 3], 3
// Expected: -1
func TestLinearSearch3(t *testing.T) {
	list := []int{1, 3}
	value := 3

	if position := Search(list, value); position != 1 {
		t.Errorf("Result was incorrect, got: %v, want: 1.", position)
	}
}

// Input: [5, 4, 2, 3, 1, 7, 0], 1
// Expected: 4
func TestLinearSearch4(t *testing.T) {
	list := []int{5, 4, 2, 3, 1, 7, 0}
	value := 1

	if position := Search(list, value); position != 4 {
		t.Errorf("Result was incorrect, got: %v, want: 4.", position)
	}
}
