package sqrt

import (
	"log"
	"testing"
)

func TestSqrt(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	inputs := []int{120}
	result := []int{10}
	res, index := 0, 0
	t.Run("test #1", func(t *testing.T) {
		for _, x := range inputs {
			res = MySqrt(x)
			if res != result[index] {
				t.Fail()
			}
			index += 1
		}
	})
}
