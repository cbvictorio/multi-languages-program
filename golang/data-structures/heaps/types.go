package heaps

type Heap interface {
	GetLeftChild() int
	GetRightChild() int
	GetParent() int
	IsEmpty() bool
	Insert(n int)
	ExtractMax() int
	GetMax() int
	Delete(index int)
	Heapify() []int
	heapifyUp(index int)
	heapifyDown(index int)
}
