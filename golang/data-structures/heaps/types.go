package heaps

type Heap interface {
	GetLeftChild(index int) int
	GetRightChild(index int) int
	GetParent(index int) int
	IsEmpty() bool
	Insert(n int)
	ExtractMax() int
	GetMax() int
	Delete(index int)
	Heapify() []int
	heapifyUp(index int)
	heapifyDown(index int)
}
