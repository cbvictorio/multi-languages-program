package trees

type Tree interface {
	IsEmpty() bool
	Size() int
	Insert(data interface{})
	InOrder()
	Height() int
}

type BasicTreeNode struct {
	Value interface{}
	Left  *BasicTreeNode
	Right *BasicTreeNode
}

type BasicBinaryTree struct {
	Root *BasicTreeNode
}
