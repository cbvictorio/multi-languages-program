package linkedlists

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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

func listToArray(list *SinglyLinkedList) (nodesArray []interface{}) {

	if !list.IsEmpty() {
		node := list.Root

		for node != nil {
			nodeValue := node.Value
			nodesArray = append(nodesArray, nodeValue)
			node = node.Next
		}
	}

	return
}

func reverseArray(array []interface{}) (reversedArray []interface{}) {
	if len(array) > 0 {
		for i := len(array) - 1; i >= 0; i-- {
			reversedArray = append(reversedArray, array[i])
		}
	}

	return
}

func AddTwoNumbers(l1 *SinglyLinkedList, l2 *SinglyLinkedList) (returnedList SinglyLinkedList) {
	totalSum := 0
	l1ToArray := listToArray(l1)
	l2ToArray := listToArray(l2)

	var l1StringSequence strings.Builder
	var l2StringSequence strings.Builder

	for i := len(l1ToArray) - 1; i >= 0; i-- {
		nodeValueToNumber := l1ToArray[i].(int)
		numberToString := strconv.Itoa(nodeValueToNumber)
		l1StringSequence.WriteString(numberToString)
	}

	for i := len(l2ToArray) - 1; i >= 0; i-- {
		nodeValueToNumber := l2ToArray[i].(int)
		numberToString := strconv.Itoa(nodeValueToNumber)
		l2StringSequence.WriteString(numberToString)
	}

	reversedL1Value, err := strconv.Atoi(l1StringSequence.String())
	if err != nil {
		log.Fatal("could not convert string sequence into number")
	}

	reversedL2Value, err := strconv.Atoi(l2StringSequence.String())
	if err != nil {
		log.Fatal("could not convert string sequence into number")
	}

	totalSum = reversedL1Value + reversedL2Value
	totalSumString := strconv.Itoa(totalSum)

	for i := 0; i <= len(totalSumString)-1; i++ {
		charToNumber, err := strconv.Atoi(string(totalSumString[i]))
		if err != nil {
			log.Fatal("could not convert a string character into number")
		}

		returnedList.Insert(charToNumber)
	}

	return
}
