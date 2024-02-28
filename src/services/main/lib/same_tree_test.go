package lib

import (
	"log"
	"testing"
)

func TestSameTree(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	// values1 := []int{100, -20, -50, -15, -60, 50, 60, 55, 85, 15, 5, -10}
	// values2 := []int{100, -20, -50, -15, -60, 50, 60, 55, 85, 15, 5, -11}
	values1 := []int{1}
	values2 := []int{0}
	root1 := NewBST(values1)
	root2 := NewBST(values2)
	i := 1
	for ; i < len(values1); i += 1 {
		root1 = root1.Insert(values1[i])
	}
	i = 1
	for ; i < len(values2); i += 1 {
		root2 = root2.Insert(values2[i])
	}
	var cursor1 *TreeNode = root1
	var cursor2 *TreeNode = root2

	result := IsSameTree(cursor1, cursor2)
	if result {
		t.Fail()
	}

}
