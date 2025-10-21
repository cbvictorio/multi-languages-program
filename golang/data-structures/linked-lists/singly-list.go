package linkedlists

import "fmt"

type SinglyLinkedList struct {
	Root *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Add to the end
func (linkedList *SinglyLinkedList) Append(data interface{}) {
	if linkedList.IsEmpty() {
		newNode := NewNode(data)
		linkedList.Root = newNode
		return
	}

	node := linkedList.Root

	for node != nil {
		if node.Next == nil {
			newNode := NewNode(data)
			node.Next = newNode
			break

		}

		node = node.Next
	}
}

// Add to the beggining
func (linkedList *SinglyLinkedList) Insert(data interface{}) {
	if linkedList.IsEmpty() {
		newNode := NewNode(data)
		linkedList.Root = newNode
		return
	}

	newNode := NewNode(data)
	rootNode := linkedList.Root
	newNode.Next = rootNode
	linkedList.Root = newNode
}

func (linkedList *SinglyLinkedList) IsEmpty() bool {
	return linkedList.Root == nil
}

func (linkedList *SinglyLinkedList) Display() {
	if linkedList.IsEmpty() {
		fmt.Print("[]")
		return
	}

	node := linkedList.Root

	for node != nil {
		fmt.Printf("[%d] ", node.Value)
		node = node.Next
	}
}

func (linkedList *SinglyLinkedList) Size() int {
	totalNodes := 0

	if linkedList.IsEmpty() {
		return 0
	}

	node := linkedList.Root
	for node != nil {
		totalNodes += 1
		node = node.Next
	}

	return totalNodes
}

func (linkedList *SinglyLinkedList) Pop() interface{} {
	if linkedList.IsEmpty() {
		return nil
	}

	if linkedList.Size() == 1 {
		value := linkedList.Root.Value
		linkedList.Root = nil
		return value
	}

	nodeA := linkedList.Root
	nodeB := nodeA.Next
	var value interface{}

	for nodeA != nil {
		if nodeB.Next == nil {
			value = nodeB.Value
			nodeA.Next = nil
			return value
		}

		nodeA = nodeB
		nodeB = nodeB.Next
	}

	return value
}

func (linkedList *SinglyLinkedList) Clear() {
	if linkedList.IsEmpty() {
		return
	}

	linkedList.Root = nil
}
