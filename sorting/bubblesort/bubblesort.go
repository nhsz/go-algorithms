package bubblesort

// Sort implements the Bubble Sort algorithm with a slice of integers
// Complexity: O(n^2)
func Sort(list []int) {
	n := len(list)
	swapCounter := -1

	for swapCounter != 0 {
		swapCounter = 0
		for i := 0; i < n; i++ {
			pass(list, i, swapCounter)
		}
	}
}

func pass(list []int, previousPasses, swapCounter int) {
	n := len(list)
	for i := 0; i+1 < n-previousPasses; i++ {
		if list[i] > list[i+1] {
			swap(list, i)
			swapCounter++
		}
	}
}

func swap(list []int, i int) {
	list[i], list[i+1] = list[i+1], list[i]
}
