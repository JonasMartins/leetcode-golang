package mediantwosortedarrays

import (
	"log"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input1 := []int{-3, -2}
	input2 := []int{10}
	result := FindMedianSortedArrays(input1, input2)
	if result == 0 {
		t.Fail()
	}
}
