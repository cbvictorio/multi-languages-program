package queues

type Queue interface {
	Print()
	Append()
	Size() int
	IsEmpty() bool
	Pop() interface{}
	PopLeft() interface{}
	Insert(data interface{})
}

type Deque struct {
	elements []interface{}
}
