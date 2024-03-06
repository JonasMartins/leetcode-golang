package linkedlistcycle

import (
	"log"
	"project/src/services/main/lib/utils"
	"testing"
)

func TestHasCycle(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{3, 2, 0, -4}
	head := utils.GenerateLinkedList(&input)
	result := HasCycle(head)
	if result {
		t.Fail()
	}
}
