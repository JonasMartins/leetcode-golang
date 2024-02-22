package lib

import (
	"math"
)

var (
	MAX_NUMS_LENGTH   = math.Pow(10.0, 4.0)
	MAX_TARGET_LENGTH = math.Pow(10.0, 9.0)
)

func TwoSum(nums []int, target int) []int {
	valid := validateTwoSumParams(&nums, &target)
	// invalid params return empty slice
	if !valid {
		return []int{}
	}

	result := []int{}
	shouldExit := false
	var j int = 0
	for i, x := range nums {
		if shouldExit {
			break
		}
		j = i + 1
		for ; j < len(nums); j += 1 {
			if x+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				shouldExit = true
				break
			}

		}
	}

	return result
}

func validateTwoSumParams(nums *[]int, target *int) bool {
	if len(*nums) < 2 || len(*nums) > int(MAX_NUMS_LENGTH) {
		return false
	}
	if *target < int(MAX_TARGET_LENGTH)*-1 || *target > int(MAX_TARGET_LENGTH) {
		return false
	}
	for _, x := range *nums {
		if x < int(MAX_TARGET_LENGTH)*-1 || x > int(MAX_TARGET_LENGTH) {
			return false
		}
	}
	return true
}
