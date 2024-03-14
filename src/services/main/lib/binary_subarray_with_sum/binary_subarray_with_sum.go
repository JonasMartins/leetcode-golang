package binarysubarraywithsum

func NumSubarraysWithSum(nums []int, goal int) int {
	result := 0
	i := goal
	if goal == 0 {
		i = 1
	}
	winPos, combinations, buff, it, size, to := 0, 0, 0, 0, 0, 0
	size = len(nums)
	prefixSum := make([]int, size)
	PrefixSumArray(&nums, &prefixSum, &size)
	for ; i <= size; i += 1 {
		result += SlideWindow(&prefixSum, &i, &goal, &winPos, &buff, &it, &combinations, &size, &to)
		winPos, combinations, it = 0, 0, 0
	}
	return result
}

func SlideWindow(prefixSumArr *[]int, wSize *int, goal *int, winPos *int, buff *int, it *int, combinations *int, size *int, to *int) int {
	// when it = 0 , buff = (*prefixSumArr)[*to]
	*to = (*winPos + *wSize) - 1
	*buff = (*prefixSumArr)[*to]
	if *buff == *goal {
		*combinations += 1
	}
	*winPos += 1
	*it = *winPos

	// avoid constant if *from = 0
	for (*winPos + *wSize) <= *size {
		*to = (*winPos + *wSize) - 1
		*buff = (*prefixSumArr)[*to] - (*prefixSumArr)[*it-1]
		if *buff == *goal {
			*combinations += 1
		}
		*winPos += 1
		*it = *winPos
	}
	return *combinations
}

func PrefixSumArray(arr *[]int, sumPrefix *[]int, size *int) {
	(*sumPrefix)[0] = (*arr)[0]
	i := 1
	for ; i < *size; i += 1 {
		(*sumPrefix)[i] = (*arr)[i] + (*sumPrefix)[i-1]
	}
}
