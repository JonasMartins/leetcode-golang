package findtreebottomleft

import (
	"log"
	"testing"
)

func TestFindTreeBottomLeftTest(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	root := NewNode(1)
	root.Left = NewNode(10)
	root.Right = NewNode(4)
	root.Left.Left = NewNode(3)
	root.Left.Left.Left = NewNode(12)
	root.Left.Left.Right = NewNode(8)
	root.Right.Left = NewNode(7)
	root.Right.Right = NewNode(9)
	root.Right.Left.Left = NewNode(6)
	root.Right.Right.Right = NewNode(2)
	result := FindBottomLeftValue(root)
	if result < 0 {
		t.Fail()
	}
}
