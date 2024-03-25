package findallduplicates

import "math"

func FindDuplicates(nums []int) []int {
	result := []int{}
	pos := 0
	// THIS WORKS ONLY BECAUSE 1 <= nums[i] <= n
	for _, x := range nums {
		pos = int(math.Abs(float64(x)) - 1)
		nums[pos] *= -1
		if nums[pos] > 0 {
			result = append(result, int(math.Abs(float64(x))))
		}
	}
	return result
}
