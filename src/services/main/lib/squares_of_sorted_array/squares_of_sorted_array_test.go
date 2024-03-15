package squaresofsortedarray

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input2 := []int{-12}
	// input := []int{
	// 	-9662, -9489, -9264, -9225, -8439, -8177, -7675, -7398, -7379,
	// 	-6374, -6295, -6199, -5457, -3899, -3762, -3696, -1638, -1316,
	// 	-341, -76, 503, 1442, 1707, 2230, 2729, 2747, 3633, 3658, 3763,
	// 	3885, 4552, 4562, 4688, 4711, 4972, 5169, 5355, 5734, 6482, 6880,
	// 	6938, 7213, 7467, 7575, 7940, 8096, 8603, 8873, 8979, 9442,
	// }
	result := SortedAquares(input2)
	if len(result) == 0 {
		t.Fail()
	}
}
