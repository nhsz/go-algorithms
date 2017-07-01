package bubblesort

// Sort implements the Bubble Sort algorithm with a slice of integers
// Complexity: O(n^2)
func Sort(list []int) {
	n := len(list)
	numberOfSwaps := -1

	// If numberOfSwaps == 0 in a pass, then the slice is already sorted
	for numberOfSwaps != 0 {
		numberOfSwaps = 0
		// Complexity: O(n^2) = n iterations * Complexity(pass) = n * n
		for i := 0; i < n; i++ {
			pass(list, i, numberOfSwaps)
		}
	}
}

// Complexity: O(n)
func pass(list []int, previousPasses, numberOfSwaps int) {
	n := len(list)
	// In the i-th iteration, the element at the n - i position of the slice
	// is placed correctly, so you don't have to check it
	for i := 0; i+1 < n-previousPasses; i++ {
		if list[i] > list[i+1] {
			swap(list, i)
			numberOfSwaps++
		}
	}
}

// Complexity: O(1)
func swap(list []int, i int) {
	list[i], list[i+1] = list[i+1], list[i]
}
