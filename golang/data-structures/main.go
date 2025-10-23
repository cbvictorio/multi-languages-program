package main

import (
	linkedlists "datastructures/linked-lists"
)

type PrintInfo struct {
	IsEmpty bool
}

func main() {
	singlyL1 := linkedlists.NewSinglyLinkedList()
	singlyL2 := linkedlists.NewSinglyLinkedList()

	singlyL1.Append(9)
	singlyL1.Append(9)
	singlyL1.Append(9)
	singlyL1.Append(9)
	singlyL1.Append(9)
	singlyL1.Append(9)
	singlyL1.Append(9)

	singlyL2.Append(9)
	singlyL2.Append(9)
	singlyL2.Append(9)
	singlyL2.Append(9)

	mergedList := linkedlists.AddTwoNumbers(singlyL1, singlyL2)
	mergedList.Display()
}
