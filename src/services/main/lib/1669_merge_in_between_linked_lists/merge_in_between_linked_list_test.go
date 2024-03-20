package mergeinbetweenlinkedlists

import (
	"log"
	"project/src/services/main/lib/utils"
	"testing"
)

func TestMergeInBetween(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	sl1 := []int{10, 1, 13, 6, 9, 5}
	sl2 := []int{1000, 1001, 1002}
	var ls1 *utils.ListNode = nil
	var ls2 *utils.ListNode = nil
	t.Run("test#1", func(t *testing.T) {
		ls1 = utils.GenerateLinkedList(&sl1)
		ls2 = utils.GenerateLinkedList(&sl2)
		ls1 = MergeInBetween(ls1, 3, 4, ls2)
		if ls1 == nil {
			t.Fail()
		}
	})
	t.Run("test#2", func(t *testing.T) {
		sl1 = []int{0, 1, 2, 3, 4, 5, 6}
		sl2 = []int{100, 101, 102, 103, 104}
		ls1 = utils.GenerateLinkedList(&sl1)
		ls2 = utils.GenerateLinkedList(&sl2)
		ls1 = MergeInBetween(ls1, 2, 5, ls2)
		if ls1 == nil {
			t.Fail()
		}
	})

}
