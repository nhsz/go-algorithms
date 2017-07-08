package binarysearch

// Search implements the Binary Search algorithm with a slice of integers
// Prerequisite: the list is already sorted
// Complexity: O(log(n))
func Search(list []int, value, min, max int) int {
	if max < min {
		return -1
	} else {
		centralPosition := (min + max) / 2

		if value < list[centralPosition] {
			return Search(list, value, min, centralPosition-1)
		} else if value > list[centralPosition] {
			return Search(list, value, centralPosition+1, max)
		}
		return centralPosition
	}
}
