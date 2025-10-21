package main

import "fmt"

func Quicksort(numbers []int) []int {
	result := []int{}
	numberOfElements := len(numbers)

	if numberOfElements <= 1 {
		return numbers
	}

	leftSideElements := []int{}
	rightSideElements := []int{}
	pivotIndex := len(numbers) - 1
	pivotValue := numbers[pivotIndex]

	for _, number := range numbers {
		if number < pivotValue {
			leftSideElements = append(leftSideElements, number)
			continue
		}

		if number > pivotValue {
			rightSideElements = append(rightSideElements, number)
			continue
		}
	}

	sortedLeftSide := Quicksort(leftSideElements)
	sortedRightSide := Quicksort(rightSideElements)

	result = append(sortedLeftSide, pivotValue)
	result = append(result, sortedRightSide...)

	return result
}

func main() {
	array := []int{3, 2, 0, 1, 2, 2, 0}
	result := Quicksort(array)
	fmt.Printf("ending result: %v", result)
}
