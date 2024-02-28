package diameterofbinarytree

import (
	"log"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	/*
		values1 := []int{1, 2, 3, 4, 5}
		root1 := NewBST(values1)
		i := 1
		for ; i < len(values1); i += 1 {
			root1 = root1.Insert(values1[i])
		}
		var cursor1 *TreeNode = root1
	*/
	cursor1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	cursor1.Left = NewNode(2)
	// cursor1.Right = NewNode(3)
	// cursor1.Left.Left = NewNode(4)
	// cursor1.Left.Right = NewNode(5)

	result := DiameterOfBinaryTree(cursor1)
	if result < 0 {
		t.Fail()
	}

}
