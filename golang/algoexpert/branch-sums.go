package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func depthFirstSearch(tree *BinaryTree, branches *[]int, totalSum int) {
	if tree != nil {
		isNodeLeaf := tree.Left == nil && tree.Right == nil
		totalSum += tree.Value

		if isNodeLeaf {
			*branches = append(*branches, totalSum)
			return
		}

		depthFirstSearch(tree.Left, branches, totalSum)
		depthFirstSearch(tree.Right, branches, totalSum)
	}
}

func BranchSums(root *BinaryTree) []int {
	var branches []int
	var totalSum int
	depthFirstSearch(root, &branches, totalSum)
	return branches
}

func main() {
	tree := &BinaryTree{Value: 1}

	// Level 1: 2 and 3
	tree.Left = &BinaryTree{Value: 2}
	tree.Right = &BinaryTree{Value: 3}

	// Level 2: 4, 5, 6, 7
	tree.Left.Left = &BinaryTree{Value: 4}
	tree.Left.Right = &BinaryTree{Value: 5}
	tree.Right.Left = &BinaryTree{Value: 6}
	tree.Right.Right = &BinaryTree{Value: 7}

	// Level 3: 8, 9, 10 (only under 4 and 5)
	tree.Left.Left.Left = &BinaryTree{Value: 8}
	tree.Left.Left.Right = &BinaryTree{Value: 9}
	tree.Left.Right.Left = &BinaryTree{Value: 10}
	result := BranchSums(tree)
	fmt.Println(result)
}
