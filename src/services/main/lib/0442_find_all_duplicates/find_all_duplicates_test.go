package findallduplicates

import (
	"log"
	"testing"
)

func TestFindAllDuplicates(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{100, 101, 105, 104, 102, 100}
	t.Run("test #1", func(t *testing.T) {
		result := FindDuplicates(input)
		if len(result) != 2 {
			t.Fail()
		}
	})
}
