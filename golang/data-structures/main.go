package main

import (
	linkedlists "datastructures/linked-lists"
	"fmt"
)

type PrintInfo struct {
	IsEmpty bool
}

func main() {
	singlyLinkedList := linkedlists.NewSinglyLinkedList()
	singlyLinkedList.Insert(9)
	singlyLinkedList.Insert(10)
	singlyLinkedList.Insert(11)
	singlyLinkedList.Append(20)
	fmt.Print("=== Linked List Values ===\n")
	singlyLinkedList.Display()
	fmt.Print("\n==========================\n")
	singlyLinkedList.Clear()

	fmt.Print("=== Linked List Values ===\n")
	singlyLinkedList.Display()
	fmt.Print("\n==========================\n")
	// poppedNode2 := singlyLinkedList.Pop()
	// fmt.Printf("\npoppedNode2 empty: %v\n", poppedNode2)
}
