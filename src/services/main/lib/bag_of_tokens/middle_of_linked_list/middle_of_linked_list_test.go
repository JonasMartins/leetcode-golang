package middleoflinkedlist

import (
	"log"
	"project/src/services/main/lib/utils"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1, 2, 3, 4, 5}
	head := utils.GenerateLinkedList(&input)
	result := MiddleNode(head)
	if result == nil {
		t.Fail()
	}
}
