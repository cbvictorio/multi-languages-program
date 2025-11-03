package main

import (
	"datastructures/trees"
	"fmt"
)

func main() {
	// Test Case 1: Sequential Ascending
	binaryTree1 := trees.NewBasicBinaryTree()
	binaryTree1.Insert(1)
	binaryTree1.Insert(2)
	binaryTree1.Insert(3)
	binaryTree1.Insert(4)
	binaryTree1.Insert(5)
	// Expected: Size: 5, InOrder: [1, 2, 3, 4, 5], Height: 4

	// Test Case 2: Sequential Descending
	binaryTree2 := trees.NewBasicBinaryTree()
	binaryTree2.Insert(50)
	binaryTree2.Insert(40)
	binaryTree2.Insert(30)
	binaryTree2.Insert(20)
	binaryTree2.Insert(10)
	// Expected: Size: 5, InOrder: [10, 20, 30, 40, 50], Height: 4

	// Test Case 3: Perfect Balanced
	binaryTree3 := trees.NewBasicBinaryTree()
	binaryTree3.Insert(50)
	binaryTree3.Insert(25)
	binaryTree3.Insert(75)
	binaryTree3.Insert(12)
	binaryTree3.Insert(37)
	binaryTree3.Insert(62)
	binaryTree3.Insert(87)
	// Expected: Size: 7, InOrder: [12, 25, 37, 50, 62, 75, 87], Height: 2

	// Test Case 4: Single Element
	binaryTree4 := trees.NewBasicBinaryTree()
	binaryTree4.Insert(100)
	// Expected: Size: 1, InOrder: [100], Height: 0

	// Test Case 5: Two Elements Left
	binaryTree5 := trees.NewBasicBinaryTree()
	binaryTree5.Insert(50)
	binaryTree5.Insert(25)
	// Expected: Size: 2, InOrder: [25, 50], Height: 1

	// Test Case 6: Two Elements Right
	binaryTree6 := trees.NewBasicBinaryTree()
	binaryTree6.Insert(50)
	binaryTree6.Insert(75)
	// Expected: Size: 2, InOrder: [50, 75], Height: 1

	// Test Case 7: Three Elements Balanced
	binaryTree7 := trees.NewBasicBinaryTree()
	binaryTree7.Insert(50)
	binaryTree7.Insert(25)
	binaryTree7.Insert(75)
	// Expected: Size: 3, InOrder: [25, 50, 75], Height: 1

	// Test Case 8: Large Numbers
	binaryTree8 := trees.NewBasicBinaryTree()
	binaryTree8.Insert(1000)
	binaryTree8.Insert(500)
	binaryTree8.Insert(1500)
	binaryTree8.Insert(250)
	binaryTree8.Insert(750)
	binaryTree8.Insert(1250)
	binaryTree8.Insert(1750)
	// Expected: Size: 7, InOrder: [250, 500, 750, 1000, 1250, 1500, 1750], Height: 2

	// Test Case 9: Negative Numbers
	binaryTree9 := trees.NewBasicBinaryTree()
	binaryTree9.Insert(0)
	binaryTree9.Insert(-10)
	binaryTree9.Insert(10)
	binaryTree9.Insert(-15)
	binaryTree9.Insert(-5)
	binaryTree9.Insert(5)
	binaryTree9.Insert(15)
	// Expected: Size: 7, InOrder: [-15, -10, -5, 0, 5, 10, 15], Height: 2

	// Test Case 10: Alternating Pattern
	binaryTree10 := trees.NewBasicBinaryTree()
	binaryTree10.Insert(50)
	binaryTree10.Insert(30)
	binaryTree10.Insert(70)
	binaryTree10.Insert(20)
	binaryTree10.Insert(60)
	binaryTree10.Insert(40)
	binaryTree10.Insert(80)
	binaryTree10.Insert(10)
	binaryTree10.Insert(35)
	// Expected: Size: 9, InOrder: [10, 20, 30, 35, 40, 50, 60, 70, 80], Height: 3

	// Test Case 11: Complete Binary Tree (15 nodes)
	binaryTree11 := trees.NewBasicBinaryTree()
	binaryTree11.Insert(8)
	binaryTree11.Insert(4)
	binaryTree11.Insert(12)
	binaryTree11.Insert(2)
	binaryTree11.Insert(6)
	binaryTree11.Insert(10)
	binaryTree11.Insert(14)
	binaryTree11.Insert(1)
	binaryTree11.Insert(3)
	binaryTree11.Insert(5)
	binaryTree11.Insert(7)
	binaryTree11.Insert(9)
	binaryTree11.Insert(11)
	binaryTree11.Insert(13)
	binaryTree11.Insert(15)
	// Expected: Size: 15, InOrder: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], Height: 3

	// Test Case 12: Fibonacci Numbers
	binaryTree12 := trees.NewBasicBinaryTree()
	binaryTree12.Insert(21)
	binaryTree12.Insert(13)
	binaryTree12.Insert(34)
	binaryTree12.Insert(8)
	binaryTree12.Insert(55)
	binaryTree12.Insert(5)
	binaryTree12.Insert(3)
	binaryTree12.Insert(2)
	binaryTree12.Insert(1)
	// Expected: Size: 9, InOrder: [1, 2, 3, 5, 8, 13, 21, 34, 55], Height: 4

	// Test Case 13: Powers of Two
	binaryTree13 := trees.NewBasicBinaryTree()
	binaryTree13.Insert(16)
	binaryTree13.Insert(8)
	binaryTree13.Insert(32)
	binaryTree13.Insert(4)
	binaryTree13.Insert(64)
	binaryTree13.Insert(2)
	binaryTree13.Insert(1)
	// Expected: Size: 7, InOrder: [1, 2, 4, 8, 16, 32, 64], Height: 3

	// Test Case 14: Mixed Positive and Negative
	binaryTree14 := trees.NewBasicBinaryTree()
	binaryTree14.Insert(0)
	binaryTree14.Insert(-50)
	binaryTree14.Insert(50)
	binaryTree14.Insert(-75)
	binaryTree14.Insert(-25)
	binaryTree14.Insert(25)
	binaryTree14.Insert(75)
	binaryTree14.Insert(-100)
	binaryTree14.Insert(100)
	// Expected: Size: 9, InOrder: [-100, -75, -50, -25, 0, 25, 50, 75, 100], Height: 3

	// Test Case 15: Prime Numbers
	binaryTree15 := trees.NewBasicBinaryTree()
	binaryTree15.Insert(23)
	binaryTree15.Insert(11)
	binaryTree15.Insert(29)
	binaryTree15.Insert(7)
	binaryTree15.Insert(19)
	binaryTree15.Insert(31)
	binaryTree15.Insert(3)
	binaryTree15.Insert(13)
	// Expected: Size: 8, InOrder: [3, 7, 11, 13, 19, 23, 29, 31], Height: 3

	// Test Case 16: Zig-Zag Pattern
	binaryTree16 := trees.NewBasicBinaryTree()
	binaryTree16.Insert(50)
	binaryTree16.Insert(25)
	binaryTree16.Insert(75)
	binaryTree16.Insert(60)
	binaryTree16.Insert(40)
	binaryTree16.Insert(85)
	binaryTree16.Insert(30)
	binaryTree16.Insert(70)
	// Expected: Size: 8, InOrder: [25, 30, 40, 50, 60, 70, 75, 85], Height: 3

	// Test Case 17: Right Skewed with Left Children
	binaryTree17 := trees.NewBasicBinaryTree()
	binaryTree17.Insert(10)
	binaryTree17.Insert(20)
	binaryTree17.Insert(5)
	binaryTree17.Insert(30)
	binaryTree17.Insert(15)
	binaryTree17.Insert(40)
	binaryTree17.Insert(25)
	// Expected: Size: 7, InOrder: [5, 10, 15, 20, 25, 30, 40], Height: 3

	// Test Case 18: Left Skewed with Right Children
	binaryTree18 := trees.NewBasicBinaryTree()
	binaryTree18.Insert(100)
	binaryTree18.Insert(80)
	binaryTree18.Insert(90)
	binaryTree18.Insert(60)
	binaryTree18.Insert(70)
	binaryTree18.Insert(40)
	binaryTree18.Insert(50)
	// Expected: Size: 7, InOrder: [40, 50, 60, 70, 80, 90, 100], Height: 3

	// Test Case 19: Small Range Values
	binaryTree19 := trees.NewBasicBinaryTree()
	binaryTree19.Insert(5)
	binaryTree19.Insert(3)
	binaryTree19.Insert(7)
	binaryTree19.Insert(2)
	binaryTree19.Insert(4)
	binaryTree19.Insert(6)
	binaryTree19.Insert(8)
	binaryTree19.Insert(1)
	binaryTree19.Insert(9)
	// Expected: Size: 9, InOrder: [1, 2, 3, 4, 5, 6, 7, 8, 9], Height: 3

	// Test Case 20: Multiples of 10
	binaryTree20 := trees.NewBasicBinaryTree()
	binaryTree20.Insert(100)
	binaryTree20.Insert(50)
	binaryTree20.Insert(150)
	binaryTree20.Insert(30)
	binaryTree20.Insert(70)
	binaryTree20.Insert(120)
	binaryTree20.Insert(180)
	binaryTree20.Insert(10)
	binaryTree20.Insert(40)
	binaryTree20.Insert(60)
	binaryTree20.Insert(90)
	// Expected: Size: 11, InOrder: [10, 30, 40, 50, 60, 70, 90, 100, 120, 150, 180], Height: 3

	// Test Case 21: Alternating Small/Large
	binaryTree21 := trees.NewBasicBinaryTree()
	binaryTree21.Insert(500)
	binaryTree21.Insert(5)
	binaryTree21.Insert(1000)
	binaryTree21.Insert(2)
	binaryTree21.Insert(250)
	binaryTree21.Insert(750)
	binaryTree21.Insert(1500)
	// Expected: Size: 7, InOrder: [2, 5, 250, 500, 750, 1000, 1500], Height: 3

	// Test Case 22: Dense Left Subtree
	binaryTree22 := trees.NewBasicBinaryTree()
	binaryTree22.Insert(100)
	binaryTree22.Insert(50)
	binaryTree22.Insert(150)
	binaryTree22.Insert(25)
	binaryTree22.Insert(75)
	binaryTree22.Insert(12)
	binaryTree22.Insert(37)
	binaryTree22.Insert(62)
	binaryTree22.Insert(87)
	binaryTree22.Insert(6)
	binaryTree22.Insert(18)
	// Expected: Size: 11, InOrder: [6, 12, 18, 25, 37, 50, 62, 75, 87, 100, 150], Height: 4

	// Test Case 23: Dense Right Subtree
	binaryTree23 := trees.NewBasicBinaryTree()
	binaryTree23.Insert(50)
	binaryTree23.Insert(25)
	binaryTree23.Insert(100)
	binaryTree23.Insert(75)
	binaryTree23.Insert(150)
	binaryTree23.Insert(62)
	binaryTree23.Insert(87)
	binaryTree23.Insert(125)
	binaryTree23.Insert(175)
	binaryTree23.Insert(112)
	binaryTree23.Insert(137)
	// Expected: Size: 11, InOrder: [25, 50, 62, 75, 87, 100, 112, 125, 137, 150, 175], Height: 4

	// Test Case 24: Reverse Sorted with Gaps
	binaryTree24 := trees.NewBasicBinaryTree()
	binaryTree24.Insert(90)
	binaryTree24.Insert(70)
	binaryTree24.Insert(50)
	binaryTree24.Insert(30)
	binaryTree24.Insert(10)
	// Expected: Size: 5, InOrder: [10, 30, 50, 70, 90], Height: 4

	// Test Case 25: Random-like Pattern 1
	binaryTree25 := trees.NewBasicBinaryTree()
	binaryTree25.Insert(45)
	binaryTree25.Insert(23)
	binaryTree25.Insert(67)
	binaryTree25.Insert(12)
	binaryTree25.Insert(34)
	binaryTree25.Insert(56)
	binaryTree25.Insert(89)
	binaryTree25.Insert(8)
	binaryTree25.Insert(28)
	binaryTree25.Insert(51)
	binaryTree25.Insert(73)
	// Expected: Size: 11, InOrder: [8, 12, 23, 28, 34, 45, 51, 56, 67, 73, 89], Height: 3

	// Test Case 26: Random-like Pattern 2
	binaryTree26 := trees.NewBasicBinaryTree()
	binaryTree26.Insert(60)
	binaryTree26.Insert(40)
	binaryTree26.Insert(80)
	binaryTree26.Insert(20)
	binaryTree26.Insert(50)
	binaryTree26.Insert(70)
	binaryTree26.Insert(90)
	binaryTree26.Insert(10)
	binaryTree26.Insert(30)
	binaryTree26.Insert(55)
	binaryTree26.Insert(65)
	binaryTree26.Insert(85)
	binaryTree26.Insert(95)
	// Expected: Size: 13, InOrder: [10, 20, 30, 40, 50, 55, 60, 65, 70, 80, 85, 90, 95], Height: 3

	// Test Case 27: Near Complete Tree
	binaryTree27 := trees.NewBasicBinaryTree()
	binaryTree27.Insert(50)
	binaryTree27.Insert(30)
	binaryTree27.Insert(70)
	binaryTree27.Insert(20)
	binaryTree27.Insert(40)
	binaryTree27.Insert(60)
	binaryTree27.Insert(80)
	binaryTree27.Insert(10)
	binaryTree27.Insert(25)
	binaryTree27.Insert(35)
	binaryTree27.Insert(45)
	// Expected: Size: 11, InOrder: [10, 20, 25, 30, 35, 40, 45, 50, 60, 70, 80], Height: 3

	// Test Case 28: Sparse Tree
	binaryTree28 := trees.NewBasicBinaryTree()
	binaryTree28.Insert(1000)
	binaryTree28.Insert(100)
	binaryTree28.Insert(10000)
	binaryTree28.Insert(10)
	binaryTree28.Insert(100000)
	binaryTree28.Insert(1)
	// Expected: Size: 6, InOrder: [1, 10, 100, 1000, 10000, 100000], Height: 3

	// Test Case 29: Binary Search Optimal
	binaryTree29 := trees.NewBasicBinaryTree()
	binaryTree29.Insert(64)
	binaryTree29.Insert(32)
	binaryTree29.Insert(96)
	binaryTree29.Insert(16)
	binaryTree29.Insert(48)
	binaryTree29.Insert(80)
	binaryTree29.Insert(112)
	binaryTree29.Insert(8)
	binaryTree29.Insert(24)
	binaryTree29.Insert(40)
	binaryTree29.Insert(56)
	binaryTree29.Insert(72)
	binaryTree29.Insert(88)
	binaryTree29.Insert(104)
	binaryTree29.Insert(120)
	// Expected: Size: 15, InOrder: [8, 16, 24, 32, 40, 48, 56, 64, 72, 80, 88, 96, 104, 112, 120], Height: 3

	// Test Case 30: Mixed Everything
	binaryTree30 := trees.NewBasicBinaryTree()
	binaryTree30.Insert(0)
	binaryTree30.Insert(-30)
	binaryTree30.Insert(30)
	binaryTree30.Insert(-50)
	binaryTree30.Insert(-10)
	binaryTree30.Insert(10)
	binaryTree30.Insert(50)
	binaryTree30.Insert(-60)
	binaryTree30.Insert(-40)
	binaryTree30.Insert(-20)
	binaryTree30.Insert(-5)
	binaryTree30.Insert(5)
	binaryTree30.Insert(20)
	binaryTree30.Insert(40)
	binaryTree30.Insert(60)
	// Expected: Size: 15, InOrder: [-60, -50, -40, -30, -20, -10, -5, 0, 5, 10, 20, 30, 40, 50, 60], Height: 3

	//Test Case 1: binaryTree1
	fmt.Print("\n == Basic Binary Tree 1 ==")
	treeSize1 := binaryTree1.Size()
	fmt.Printf("\nSize(): %d", treeSize1)

	treeInOrderElements1 := binaryTree1.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements1)

	treeHeight1 := binaryTree1.Height()
	fmt.Printf("\nHeight(): %d", treeHeight1)

	// Test Case 2: binaryTree2
	fmt.Print("\n\n == Basic Binary Tree 2 ==")
	treeSize2 := binaryTree2.Size()
	fmt.Printf("\nSize(): %d", treeSize2)

	treeInOrderElements2 := binaryTree2.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements2)

	treeHeight2 := binaryTree2.Height()
	fmt.Printf("\nHeight(): %d", treeHeight2)

	// Test Case 3: binaryTree3
	fmt.Print("\n\n == Basic Binary Tree 3 ==")
	treeSize3 := binaryTree3.Size()
	fmt.Printf("\nSize(): %d", treeSize3)

	treeInOrderElements3 := binaryTree3.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements3)

	treeHeight3 := binaryTree3.Height()
	fmt.Printf("\nHeight(): %d", treeHeight3)

	// Test Case 4: binaryTree4
	fmt.Print("\n\n == Basic Binary Tree 4 ==")
	treeSize4 := binaryTree4.Size()
	fmt.Printf("\nSize(): %d", treeSize4)

	treeInOrderElements4 := binaryTree4.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements4)

	treeHeight4 := binaryTree4.Height()
	fmt.Printf("\nHeight(): %d", treeHeight4)

	// Test Case 5: binaryTree5
	fmt.Print("\n\n == Basic Binary Tree 5 ==")
	treeSize5 := binaryTree5.Size()
	fmt.Printf("\nSize(): %d", treeSize5)

	treeInOrderElements5 := binaryTree5.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements5)

	treeHeight5 := binaryTree5.Height()
	fmt.Printf("\nHeight(): %d", treeHeight5)

	// Test Case 6: binaryTree6
	fmt.Print("\n\n == Basic Binary Tree 6 ==")
	treeSize6 := binaryTree6.Size()
	fmt.Printf("\nSize(): %d", treeSize6)

	treeInOrderElements6 := binaryTree6.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements6)

	treeHeight6 := binaryTree6.Height()
	fmt.Printf("\nHeight(): %d", treeHeight6)

	// Test Case 7: binaryTree7
	fmt.Print("\n\n == Basic Binary Tree 7 ==")
	treeSize7 := binaryTree7.Size()
	fmt.Printf("\nSize(): %d", treeSize7)

	treeInOrderElements7 := binaryTree7.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements7)

	treeHeight7 := binaryTree7.Height()
	fmt.Printf("\nHeight(): %d", treeHeight7)

	// Test Case 8: binaryTree8
	fmt.Print("\n\n == Basic Binary Tree 8 ==")
	treeSize8 := binaryTree8.Size()
	fmt.Printf("\nSize(): %d", treeSize8)

	treeInOrderElements8 := binaryTree8.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements8)

	treeHeight8 := binaryTree8.Height()
	fmt.Printf("\nHeight(): %d", treeHeight8)

	// Test Case 9: binaryTree9
	fmt.Print("\n\n == Basic Binary Tree 9 ==")
	treeSize9 := binaryTree9.Size()
	fmt.Printf("\nSize(): %d", treeSize9)

	treeInOrderElements9 := binaryTree9.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements9)

	treeHeight9 := binaryTree9.Height()
	fmt.Printf("\nHeight(): %d", treeHeight9)

	// Test Case 10: binaryTree10
	fmt.Print("\n\n == Basic Binary Tree 10 ==")
	treeSize10 := binaryTree10.Size()
	fmt.Printf("\nSize(): %d", treeSize10)

	treeInOrderElements10 := binaryTree10.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements10)

	treeHeight10 := binaryTree10.Height()
	fmt.Printf("\nHeight(): %d", treeHeight10)

	// Test Case 11: binaryTree11
	fmt.Print("\n\n == Basic Binary Tree 11 ==")
	treeSize11 := binaryTree11.Size()
	fmt.Printf("\nSize(): %d", treeSize11)

	treeInOrderElements11 := binaryTree11.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements11)

	treeHeight11 := binaryTree11.Height()
	fmt.Printf("\nHeight(): %d", treeHeight11)

	// Test Case 12: binaryTree12
	fmt.Print("\n\n == Basic Binary Tree 12 ==")
	treeSize12 := binaryTree12.Size()
	fmt.Printf("\nSize(): %d", treeSize12)

	treeInOrderElements12 := binaryTree12.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements12)

	treeHeight12 := binaryTree12.Height()
	fmt.Printf("\nHeight(): %d", treeHeight12)

	// Test Case 13: binaryTree13
	fmt.Print("\n\n == Basic Binary Tree 13 ==")
	treeSize13 := binaryTree13.Size()
	fmt.Printf("\nSize(): %d", treeSize13)

	treeInOrderElements13 := binaryTree13.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements13)

	treeHeight13 := binaryTree13.Height()
	fmt.Printf("\nHeight(): %d", treeHeight13)

	// Test Case 14: binaryTree14
	fmt.Print("\n\n == Basic Binary Tree 14 ==")
	treeSize14 := binaryTree14.Size()
	fmt.Printf("\nSize(): %d", treeSize14)

	treeInOrderElements14 := binaryTree14.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements14)

	treeHeight14 := binaryTree14.Height()
	fmt.Printf("\nHeight(): %d", treeHeight14)

	// Test Case 15: binaryTree15
	fmt.Print("\n\n == Basic Binary Tree 15 ==")
	treeSize15 := binaryTree15.Size()
	fmt.Printf("\nSize(): %d", treeSize15)

	treeInOrderElements15 := binaryTree15.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements15)

	treeHeight15 := binaryTree15.Height()
	fmt.Printf("\nHeight(): %d", treeHeight15)

	// Test Case 16: binaryTree16
	fmt.Print("\n\n == Basic Binary Tree 16 ==")
	treeSize16 := binaryTree16.Size()
	fmt.Printf("\nSize(): %d", treeSize16)

	treeInOrderElements16 := binaryTree16.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements16)

	treeHeight16 := binaryTree16.Height()
	fmt.Printf("\nHeight(): %d", treeHeight16)

	// Test Case 17: binaryTree17
	fmt.Print("\n\n == Basic Binary Tree 17 ==")
	treeSize17 := binaryTree17.Size()
	fmt.Printf("\nSize(): %d", treeSize17)

	treeInOrderElements17 := binaryTree17.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements17)

	treeHeight17 := binaryTree17.Height()
	fmt.Printf("\nHeight(): %d", treeHeight17)

	// Test Case 18: binaryTree18
	fmt.Print("\n\n == Basic Binary Tree 18 ==")
	treeSize18 := binaryTree18.Size()
	fmt.Printf("\nSize(): %d", treeSize18)

	treeInOrderElements18 := binaryTree18.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements18)

	treeHeight18 := binaryTree18.Height()
	fmt.Printf("\nHeight(): %d", treeHeight18)

	// Test Case 19: binaryTree19
	fmt.Print("\n\n == Basic Binary Tree 19 ==")
	treeSize19 := binaryTree19.Size()
	fmt.Printf("\nSize(): %d", treeSize19)

	treeInOrderElements19 := binaryTree19.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements19)

	treeHeight19 := binaryTree19.Height()
	fmt.Printf("\nHeight(): %d", treeHeight19)

	// Test Case 20: binaryTree20
	fmt.Print("\n\n == Basic Binary Tree 20 ==")
	treeSize20 := binaryTree20.Size()
	fmt.Printf("\nSize(): %d", treeSize20)

	treeInOrderElements20 := binaryTree20.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements20)

	treeHeight20 := binaryTree20.Height()
	fmt.Printf("\nHeight(): %d", treeHeight20)

	// Test Case 21: binaryTree21
	fmt.Print("\n\n == Basic Binary Tree 21 ==")
	treeSize21 := binaryTree21.Size()
	fmt.Printf("\nSize(): %d", treeSize21)

	treeInOrderElements21 := binaryTree21.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements21)

	treeHeight21 := binaryTree21.Height()
	fmt.Printf("\nHeight(): %d", treeHeight21)

	// Test Case 22: binaryTree22
	fmt.Print("\n\n == Basic Binary Tree 22 ==")
	treeSize22 := binaryTree22.Size()
	fmt.Printf("\nSize(): %d", treeSize22)

	treeInOrderElements22 := binaryTree22.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements22)

	treeHeight22 := binaryTree22.Height()
	fmt.Printf("\nHeight(): %d", treeHeight22)

	// Test Case 23: binaryTree23
	fmt.Print("\n\n == Basic Binary Tree 23 ==")
	treeSize23 := binaryTree23.Size()
	fmt.Printf("\nSize(): %d", treeSize23)

	treeInOrderElements23 := binaryTree23.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements23)

	treeHeight23 := binaryTree23.Height()
	fmt.Printf("\nHeight(): %d", treeHeight23)

	// Test Case 24: binaryTree24
	fmt.Print("\n\n == Basic Binary Tree 24 ==")
	treeSize24 := binaryTree24.Size()
	fmt.Printf("\nSize(): %d", treeSize24)

	treeInOrderElements24 := binaryTree24.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements24)

	treeHeight24 := binaryTree24.Height()
	fmt.Printf("\nHeight(): %d", treeHeight24)

	// Test Case 25: binaryTree25
	fmt.Print("\n\n == Basic Binary Tree 25 ==")
	treeSize25 := binaryTree25.Size()
	fmt.Printf("\nSize(): %d", treeSize25)

	treeInOrderElements25 := binaryTree25.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements25)

	treeHeight25 := binaryTree25.Height()
	fmt.Printf("\nHeight(): %d", treeHeight25)

	// Test Case 26: binaryTree26
	fmt.Print("\n\n == Basic Binary Tree 26 ==")
	treeSize26 := binaryTree26.Size()
	fmt.Printf("\nSize(): %d", treeSize26)

	treeInOrderElements26 := binaryTree26.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements26)

	treeHeight26 := binaryTree26.Height()
	fmt.Printf("\nHeight(): %d", treeHeight26)

	// Test Case 27: binaryTree27
	fmt.Print("\n\n == Basic Binary Tree 27 ==")
	treeSize27 := binaryTree27.Size()
	fmt.Printf("\nSize(): %d", treeSize27)

	treeInOrderElements27 := binaryTree27.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements27)

	treeHeight27 := binaryTree27.Height()
	fmt.Printf("\nHeight(): %d", treeHeight27)

	// Test Case 28: binaryTree28
	fmt.Print("\n\n == Basic Binary Tree 28 ==")
	treeSize28 := binaryTree28.Size()
	fmt.Printf("\nSize(): %d", treeSize28)

	treeInOrderElements28 := binaryTree28.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements28)

	treeHeight28 := binaryTree28.Height()
	fmt.Printf("\nHeight(): %d", treeHeight28)

	// Test Case 29: binaryTree29
	fmt.Print("\n\n == Basic Binary Tree 29 ==")
	treeSize29 := binaryTree29.Size()
	fmt.Printf("\nSize(): %d", treeSize29)

	treeInOrderElements29 := binaryTree29.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements29)

	treeHeight29 := binaryTree29.Height()
	fmt.Printf("\nHeight(): %d", treeHeight29)

	// Test Case 30: binaryTree30
	fmt.Print("\n\n == Basic Binary Tree 30 ==")
	treeSize30 := binaryTree30.Size()
	fmt.Printf("\nSize(): %d", treeSize30)

	treeInOrderElements30 := binaryTree30.InOrder()
	fmt.Printf("\nInOrder(): %s", treeInOrderElements30)

	treeHeight30 := binaryTree30.Height()
	fmt.Printf("\nHeight(): %d", treeHeight30)
}
