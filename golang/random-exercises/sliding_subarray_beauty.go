package main

import (
	"fmt"
	"sort"
)

/*
	Given an integer array nums containing n integers, find the beauty of each subarray of size k.

	The beauty of a subarray is the xth smallest integer in the subarray if it is negative, or 0 if there are fewer than x negative integers.

	Return an integer array containing n - k + 1 integers, which denote the beauty of the subarrays in order from the first index in the array.

	A subarray is a contiguous non-empty sequence of elements within an array.
*/

func GetSubarrayBeauty(nums []int, k int, x int) []int {
	var subArray []int

	windowLimit := len(nums) - k

	for i := 0; i <= windowLimit; i++ {
		lastWindowIndex := i + k
		windowArray := append([]int{}, nums[i:lastWindowIndex]...)
		sort.Ints(windowArray)
		elementToAdd := windowArray[x-1]

		if elementToAdd >= 0 {
			subArray = append(subArray, 0)
		} else {
			subArray = append(subArray, elementToAdd)
		}

	}

	return subArray
}

func main() {
	elements := []int{1, -1, -3, -2, 3}
	subArray := GetSubarrayBeauty(elements, 3, 2)

	elements2 := []int{-1, -2, -3, -4, -5}
	subArray2 := GetSubarrayBeauty(elements2, 2, 2)

	elements3 := []int{-3, 1, 2, -3, 0, -3}
	subArray3 := GetSubarrayBeauty(elements3, 2, 1)

	fmt.Printf("subArray: %v\n", subArray)
	fmt.Printf("subArray2: %v\n", subArray2)
	fmt.Printf("subArray3: %v\n", subArray3)
}
