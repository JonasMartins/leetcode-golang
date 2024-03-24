package findduplicatenumber

import (
	"log"
	"testing"
)

func TestFindDuplicateNumber(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1, 2, 6, 7, 9, 8, 4, 5, 5, 3}
	t.Run("test #1", func(t *testing.T) {
		result := FindDuplicate(input)
		if result < 1 {
			t.Fail()
		}
	})
}
