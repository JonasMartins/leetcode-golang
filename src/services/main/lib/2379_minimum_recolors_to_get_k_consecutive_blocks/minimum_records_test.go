package minimumrecolorstogetkconsecutiveblocks

import (
	"log"
	"testing"
)

func TestMinumunRecolors(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	inputs := map[string]int{
		"WBBWWBBWBW": 4,
	}
	result := []int{1}
	res, index := 0, 0
	t.Run("test #1", func(t *testing.T) {
		for x, i := range inputs {
			res = MinimumRecolors(x, i)
			if res != result[index] {
				t.Fail()
			}
			index += 1
		}
	})
}
