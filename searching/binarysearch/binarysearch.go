package binarysearch

// RecursiveSearch implements a recursive version of the Binary Search
// algorithm with a slice of integers
// Prerequisite: the list is already sorted
// Complexity: O(log(n))
func RecursiveSearch(list []int, value, min, max int) int {
	if max < min {
		return -1
	}

	mid := (min + max) / 2

	if value < list[mid] {
		return RecursiveSearch(list, value, min, mid-1)
	} else if value > list[mid] {
		return RecursiveSearch(list, value, mid+1, max)
	}
	return mid
}

// IterativeSearch implements an iterative version of the Binary Search
// algorithm with a slice of integers
// Prerequisite: the list is already sorted
// Complexity: O(log(n))
func IterativeSearch(list []int, value int) int {
	min := 0
	max := len(list) - 1

	for min <= max {
		mid := (min + max) / 2
		if value == list[mid] {
			return mid
		} else if value < list[mid] {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
