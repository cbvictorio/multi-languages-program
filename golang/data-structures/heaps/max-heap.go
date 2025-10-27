package heaps

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (heap *MaxHeap) GetLeftChild(index int) int {
	return (2 * index) + 1
}

func (heap *MaxHeap) GetRightChild(index int) int {
	return (2 * index) + 2
}

func (heap *MaxHeap) GetParent(index int) int {
	return (index - 1) / 2
}

func (heap *MaxHeap) IsEmpty() bool {
	return len(heap.array) == 0
}

func (heap *MaxHeap) Size() int {
	return len(heap.array)
}

func (heap *MaxHeap) GetMax() int {
	if heap.IsEmpty() {
		return -1
	}

	return heap.array[0]
}

func (heap *MaxHeap) PrintValues() {
	if heap.IsEmpty() {
		fmt.Println("The heap is empty")
		return
	}

	fmt.Print("\n")
	for _, value := range heap.array {
		fmt.Printf("[%d] ", value)
	}
	fmt.Print("\n")
}

func (heap *MaxHeap) hasElementLeftChild(index int) bool {
	leftChildIndex := heap.GetLeftChild(index)
	return len(heap.array) > leftChildIndex
}

func (heap *MaxHeap) hasElementRightChild(index int) bool {
	rightChildIndex := heap.GetRightChild(index)
	return len(heap.array) > rightChildIndex
}

func (heap *MaxHeap) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := heap.GetParent(index)
	parentValue := heap.array[parentIndex]
	currentValue := heap.array[index]

	if currentValue > parentValue {
		heap.array[parentIndex] = currentValue
		heap.array[index] = parentValue
		heap.heapifyUp(parentIndex)
	}
}

func (heap *MaxHeap) heapifyDown(index int) {
	/*
		if the element we are trying to heapify down is the last one,
		there's no point of continue checking down
	*/
	if index == len(heap.array)-1 {
		fmt.Print("the last element has been reached")
		return
	}

	elementValue := heap.array[index]

	// left side
	if heap.hasElementLeftChild(index) {
		leftChildIndex := heap.GetLeftChild(index)
		leftChildValue := heap.array[leftChildIndex]

		if leftChildValue > elementValue {
			heap.array[index] = leftChildValue
			heap.array[leftChildIndex] = elementValue
			heap.heapifyDown(leftChildIndex)
			return
		}
	}

	// right side
	if heap.hasElementRightChild(index) {
		rightChildIndex := heap.GetRightChild(index)
		rightChildValue := heap.array[rightChildIndex]

		if rightChildValue > elementValue {
			heap.array[index] = rightChildValue
			heap.array[rightChildIndex] = elementValue
			heap.heapifyDown(rightChildIndex)
			return
		}
	}
}

func (heap *MaxHeap) Insert(number int) {
	heap.array = append(heap.array, number)

	if heap.Size() < 2 {
		return
	}

	recentlyAddedIndex := len(heap.array) - 1
	heap.heapifyUp(recentlyAddedIndex)
}

func (heap *MaxHeap) ExtractMax() int {
	if heap.IsEmpty() {
		fmt.Print("\nthe heap is empty\n")
		return -1
	}

	// grab last element index and value
	lastElementIndex := len(heap.array) - 1
	lastElementValue := heap.array[lastElementIndex]

	// save the first value and swap with the last element
	maxValue := heap.array[0]
	heap.array[0] = lastElementValue

	// remove last element
	heap.array = heap.array[:lastElementIndex]

	// run the heapify down to check the integrity of the heap
	heap.heapifyDown(0)

	return maxValue
}
