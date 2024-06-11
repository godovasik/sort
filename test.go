package main

import "fmt"

func main() {
	arr := make([]int, 10)
	brr := make([]int, 10)

	fmt.Println("Initial arrays:")
	fmt.Println("arr:", arr)
	fmt.Println("brr:", brr)

	copy(brr, arr) // brr is now a separate array with the same elements as arr

	// Modify arr
	arr[0] = 42

	fmt.Println("\nAfter modifying arr:")
	fmt.Println("arr:", arr)
	fmt.Println("brr:", brr)
}
