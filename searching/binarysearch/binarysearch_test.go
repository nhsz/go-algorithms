package binarysearch

import (
	"testing"
)

// Input: [], 3
// Expected: -1
func TestBinarySearch1(t *testing.T) {
	list := []int{}
	n := len(list) - 1
	value := 3

	if position := Search(list, value, 0, n); position != -1 {
		t.Errorf("Result was incorrect, got: %v, want: -1.", position)
	}
}

// Input: [1], 3
// Expected: -1
func TestBinarySearch2(t *testing.T) {
	list := []int{1}
	n := len(list) - 1
	value := 3

	if position := Search(list, value, 0, n); position != -1 {
		t.Errorf("Result was incorrect, got: %v, want: -1.", position)
	}
}

// Input: [1, 3], 3
// Expected: -1
func TestBinarySearch3(t *testing.T) {
	list := []int{1, 3}
	n := len(list) - 1
	value := 3

	if position := Search(list, value, 0, n); position != 1 {
		t.Errorf("Result was incorrect, got: %v, want: 1.", position)
	}
}

// Input: [0, 1, 2, 3, 4, 5, 7], 4
// Expected: 4
func TestBinarySearch4(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 7}
	n := len(list) - 1
	value := 4

	if position := Search(list, value, 0, n); position != 4 {
		t.Errorf("Result was incorrect, got: %v, want: 4.", position)
	}
}
