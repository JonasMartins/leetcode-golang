package minimumcommomvalue

import (
	"log"
	"testing"
)

func TestGetCommon(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4}
	result := 0
	t.Run("test #1", func(t *testing.T) {
		result = GetCommon(input1, input2)
		if result != 2 {
			t.Fail()
		}
	})
	t.Run("test #2", func(t *testing.T) {
		input1 = []int{1, 2, 3, 6}
		input2 = []int{2, 3, 4, 5}
		result = GetCommon(input1, input2)
		if result != 2 {
			t.Fail()
		}
	})
}
