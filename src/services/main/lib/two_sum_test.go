package lib

import (
	"log"
	"math"
	"testing"
)

func TestTwoSum(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	t.Run("Should get invalid target", func(t *testing.T) {
		nums, target := []int{2}, 9
		result := TwoSum(nums, target)
		// should return a empty slice
		if len(result) != 0 {
			t.Fail()
		}
	})

	t.Run("Should get invalid max num element", func(t *testing.T) {
		maxF := math.Pow(10.0, 9.0)
		maxI := int(maxF) + 1
		nums, target := []int{1, 2, 3, maxI}, 3
		result := TwoSum(nums, target)
		// should return a empty slice
		if len(result) != 0 {
			t.Fail()
		}
	})

	t.Run("Should get invalid min num element", func(t *testing.T) {
		maxF := math.Pow(10.0, 9.0)
		minI := (int(maxF) + 1) * -1
		nums, target := []int{1, 2, 3, minI}, 3
		result := TwoSum(nums, target)
		// should return a empty slice
		if len(result) != 0 {
			t.Fail()
		}
	})

	t.Run("Should get invalid max target", func(t *testing.T) {
		maxF := math.Pow(10.0, 9.0)
		target := (int(maxF) + 1) * -1
		nums := []int{1, 2, 3}
		result := TwoSum(nums, target)
		// should return a empty slice
		if len(result) != 0 {
			t.Fail()
		}
	})

	t.Run("Should get a valid result", func(t *testing.T) {
		nums, target := []int{13, 23, 17, 19, 7, 5, 65, 44, 1}, 8
		result := TwoSum(nums, target)
		// should not return a empty slice
		if len(result) == 0 {
			t.Fail()
		} else if len(result) == 2 {
			if nums[result[0]]+nums[result[1]] != target {
				t.Fail()
			}
		}

	})
}
