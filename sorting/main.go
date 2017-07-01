package main

import (
	"./bubblesort"
	"fmt"
)

func main() {
	fmt.Println("SORTING ALGORITHMS")
	fmt.Println()

	list := []int{5, 4, 2, 3, 1, 7, 0}
	fmt.Println("Original list: ", list)
	fmt.Println()

	bubblesort.Sort(list)
	fmt.Println("Bubble Sort: ", list)
}
