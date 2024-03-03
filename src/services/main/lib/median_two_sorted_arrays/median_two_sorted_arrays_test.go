package mediantwosortedarrays

import (
	"log"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input1 := []int{1}
	input2 := []int{2}
	result := FindMedianSortedArrays(input1, input2)
	if result == 0 {
		t.Fail()
	}
}
