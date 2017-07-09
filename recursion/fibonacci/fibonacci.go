package fibonacci

// RecursiveFib implements a recursive version of the Fibonacci sequence
// Complexity:
func RecursiveFib(n int) int {
	if n <= 1 {
		return n
	}
	return RecursiveFib(n-1) + RecursiveFib(n-2)
}

// IterativeFib implements an iterative version of the Fibonacci sequence
// Complexity: O(n)
func IterativeFib(n int) int {
	if n <= 1 {
		return n
	}

	first := 0
	second := 1
	next := first + second

	for i := 2; i < n; i++ {
		first = second
		second = next
		next = first + second
	}
	return next
}
