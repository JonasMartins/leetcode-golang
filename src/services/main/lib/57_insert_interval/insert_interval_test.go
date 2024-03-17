package insertinterval

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := [][]int{
		{6, 7},
		{9, 12},
	}
	newInterval := []int{1, 1}
	result := Insert(input, newInterval)
	if len(result) == 0 {
		t.Fail()
	}
}
