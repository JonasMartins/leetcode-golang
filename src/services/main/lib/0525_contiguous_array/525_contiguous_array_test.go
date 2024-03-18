package contiguousarray

import (
	"log"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0}
	result := FindMaxLength(input)
	if result < 0 {
		t.Fail()
	}
}
