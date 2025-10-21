package linkedlists

type Node struct {
	Value interface{}
	Next  *Node
}

type NodeActions interface {
	GetNext() *Node
	SetNext()
}

type LinkedList interface {
	Insert(data int)
	Append(data int)
	Pop()
	Size() int
	IsEmpty() bool
	Display()
	Clear()
}

func NewNode(data interface{}) *Node {
	return &Node{Next: nil, Value: data}
}
