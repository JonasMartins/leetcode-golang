package evenoddtree

import (
	"log"
	"testing"
)

func TestFindEvenOddTree(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	root := NewNode(1)
	root.Left = NewNode(10)
	root.Right = NewNode(8)
	root.Left.Left = NewNode(3)
	root.Left.Left.Left = NewNode(12)
	root.Left.Left.Right = NewNode(8)
	root.Right.Left = NewNode(5)
	root.Right.Right = NewNode(11)
	root.Right.Left.Left = NewNode(6)
	root.Right.Right.Right = NewNode(2)
	result := IsEvenOddTree(root)
	if result {
		t.Fail()
	}

}
