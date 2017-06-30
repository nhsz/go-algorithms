package main

import (
	"./bubblesort"
	"fmt"
)

func main() {
	fmt.Println("SORTING ALGORITHMS\n")

	list := []int{5, 4, 2, 3, 1, 7, 0}
	fmt.Println("Original list: ", list, "\n")

	bubblesort.Sort(list)
	fmt.Println("Bubble Sort: ", list)
}
