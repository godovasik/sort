package main

import (
	"github.com/gdamore/tcell"
	"math/rand"
)

func cocktailShakerSort(arr, brr []int, screen tcell.Screen) {
	_, height := screen.Size()
	n := len(arr)
	swapped := true
	start := 0
	end := n - 1

	for swapped {
		swapped = false

		// Loop from left to right, similar to the bubble sort.
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				copy(brr, arr)
				arr[i], arr[i+1] = arr[i+1], arr[i] // Swap elements
				screen.Clear()
				draw(arr, brr, height, screen)
				swapped = true
			}
		}

		// If no numbers were swapped during the last iteration,
		// the array is already sorted, and we can terminate early.
		if !swapped {
			break
		}

		// Otherwise, reset the swapped flag to false so it can be
		// used in the next stage.
		swapped = false

		// Move the endpoint back by one, because the item at the
		// end of the array has been moved to its correct place.
		end--

		// Reverse the direction of the scan and repeat the process.
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				copy(brr, arr)
				arr[i], arr[i+1] = arr[i+1], arr[i] // Swap elements
				screen.Clear()
				draw(arr, brr, height, screen)
				swapped = true
			}
		}

		// Increase the starting point, because the item at the
		// start of the array has been moved to its correct place.
		start++
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, val := range arr[1:] {
		if val < pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func bubble(arr, brr []int, screen tcell.Screen) {
	_, height := screen.Size()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				copy(brr, arr)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				screen.Clear()
				draw(arr, brr, height, screen)
			}
		}
	}
}

func quick(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(nums)

	// Move the pivot to the right
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	// Place the pivot after the elements smaller than it
	nums[left], nums[right] = nums[right], nums[left]

	// Recurse into the two subarrays
	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}
