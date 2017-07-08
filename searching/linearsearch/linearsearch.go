package linearsearch

// Search implements the Linear Search algorithm with a slice of integers
// It returns the position of the wanted value if it's on the list.
// Otherwise, it returns -1
// Complexity: O(n)
func Search(list []int, value int) int {
	n := len(list)

	for i := 0; i < n; i++ {
		if list[i] == value {
			return i
		}
	}
	return -1
}
