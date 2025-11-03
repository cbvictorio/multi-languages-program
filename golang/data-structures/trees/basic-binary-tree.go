package trees

import (
	"datastructures/queues"
	"fmt"
	"strings"
)

/*
	Basic Binary Tree implementation is assuming we only have
	a Left and Right Children as children, not an array of children
*/

func NewBasicBinaryTree() *BasicBinaryTree {
	return &BasicBinaryTree{}
}

func NewBasicTreeNode(value interface{}) *BasicTreeNode {
	return &BasicTreeNode{Value: value}
}

func (basicBinaryTree *BasicBinaryTree) IsEmpty() bool {
	return basicBinaryTree.Root == nil
}

/*
we assume that new values will go to the left of the tree
if they are equals or lesser than the current node's value
*/
func (basicBinaryTree *BasicBinaryTree) Insert(newValue int) {
	newNode := NewBasicTreeNode(newValue)

	if basicBinaryTree.IsEmpty() {
		basicBinaryTree.Root = newNode
		return
	}

	node := basicBinaryTree.Root
	for node != nil {
		currentNodeValue, _ := node.Value.(int)

		if newValue <= currentNodeValue {
			if node.Left != nil {
				node = node.Left
				continue
			}

			node.Left = newNode
			break
		}

		if node.Right != nil {
			node = node.Right
			continue
		}

		node.Right = newNode
		break
	}
}

func inOrderTraversal(node *BasicTreeNode, sb *strings.Builder) string {
	if node != nil {
		inOrderTraversal(node.Left, sb)
		nodeValueToString := fmt.Sprintf("%v", node.Value)
		sb.WriteString("[" + nodeValueToString + "] ")
		inOrderTraversal(node.Right, sb)
	}

	return sb.String()
}

func countTraversal(node *BasicTreeNode, sum *int) {
	if node != nil {
		countTraversal(node.Left, sum)
		*sum++
		countTraversal(node.Right, sum)
	}
}

func (basicBinaryTree *BasicBinaryTree) InOrder() string {
	if basicBinaryTree.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder

	elements := inOrderTraversal(basicBinaryTree.Root, &sb)
	return elements
}

func (basicBinaryTree *BasicBinaryTree) Size() int {
	sum := 0

	if !basicBinaryTree.IsEmpty() {
		rootNode := basicBinaryTree.Root
		countTraversal(rootNode, &sum)
	}

	return sum
}

func (basicBinaryTree *BasicBinaryTree) Height() int {

	if basicBinaryTree.IsEmpty() {
		return -1
	}

	// single element case
	rootNode := basicBinaryTree.Root
	if rootNode.Left == nil && rootNode.Right == nil {
		return 0
	}

	queue := queues.NewDeque()
	queue.Insert(basicBinaryTree.Root)
	treeHeight := -1

	for queue.Size() > 0 {
		i := 0
		queueSize := queue.Size()

		for i < queueSize {
			currentNode := queue.PopLeft().(*BasicTreeNode)

			if currentNode.Left != nil {
				queue.Append(currentNode.Left)
			}

			if currentNode.Right != nil {
				queue.Append(currentNode.Right)
			}

			i++
		}

		treeHeight++
	}

	return treeHeight

}
