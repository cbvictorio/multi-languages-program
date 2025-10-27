package main

import (
	"datastructures/heaps"
	"fmt"
)

func main() {
	maxHeap := heaps.NewMaxHeap()
	maxHeap.Insert(5)
	maxHeap.Insert(12)
	maxHeap.Insert(64)
	maxHeap.Insert(1)
	maxHeap.Insert(37)
	maxHeap.Insert(90)
	maxHeap.Insert(91)
	maxHeap.Insert(97)

	maxHeap.PrintValues()

	maxHeapValue := maxHeap.ExtractMax()
	fmt.Printf("the deleted value was: %d", maxHeapValue)
	maxHeap.PrintValues()
}
