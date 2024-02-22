package lib

import (
	"log"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	// l1 := BuildLinkedListFromArray(&[]uint8{0})
	// l2 := BuildLinkedListFromArray(&[]uint8{0})

	// l1 := BuildLinkedListFromArray(&[]uint8{2, 4, 3})
	// l2 := BuildLinkedListFromArray(&[]uint8{5, 6, 4})

	// l1 := BuildLinkedListFromArray(&[]uint8{9, 9, 9, 9, 9, 9, 9})
	// l2 := BuildLinkedListFromArray(&[]uint8{9, 9, 9, 9})

	// l1 := BuildLinkedListFromArray(&[]uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	// l2 := BuildLinkedListFromArray(&[]uint8{5, 6, 4})

	l1 := BuildLinkedListFromArray(&[]uint8{0})
	l2 := BuildLinkedListFromArray(&[]uint8{9, 7, 1})
	result := AddTwoNumbers(l1, l2)
	if result == nil {
		t.Fail()
	}
	PrintLinkedList(result)

}
