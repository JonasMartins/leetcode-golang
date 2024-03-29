package removenthfromend

import (
	"log"
	"project/src/services/main/lib/utils"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1}
	head := utils.GenerateLinkedList(&input)
	head = RemoveNthFromEnd(head, 1)
	if head == nil {
		t.Fail()
	}

	log.Println("test finished")
}
