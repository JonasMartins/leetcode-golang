package reorderlist

import (
	"log"
	"project/src/services/main/lib/utils"
	"testing"
)

func TestReorderList(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	t.Run("test #1", func(t *testing.T) {
		head := utils.GenerateLinkedList(&input)
		ReorderList(head)
		if head == nil {
			t.Fail()
		}
	})
}
