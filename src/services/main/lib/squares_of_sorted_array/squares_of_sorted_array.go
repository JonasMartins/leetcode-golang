package squaresofsortedarray

func SortedAquares(nums []int) []int {
	result := []int{}
	ptrResult := &result
	ptrNums := &nums
	last := len(nums) - 1
	if last == 0 {
		result = append(result, nums[0]*nums[0])
		return result
	}
	if nums[0] < 0 && nums[last] <= 0 {
		SquareInReverseOrder(ptrNums, ptrResult)
	}
	if nums[0] >= 0 {
		SquareInNormalOrder(ptrNums, ptrResult)
	}
	if nums[0] < 0 && nums[last] > 0 {
		SquareInNegativeToPositiveOrder(ptrNums, ptrResult)
	}

	return result
}

func SquareInNegativeToPositiveOrder(nums *[]int, result *[]int) {
	last := len(*nums) - 1
	i := 0
	backIter, forwardIter := 1, 1
	hasRoomBack, hasRoomForward := false, false
	backBuff, forwardBuff := 0, 0
	for {
		if i == last || (*nums)[i] >= 0 {
			break
		}
		i += 1
	}
	backIter = i - 1
	forwardIter = i
	for {
		hasRoomBack = backIter >= 0
		hasRoomForward = forwardIter <= last
		if !hasRoomBack && !hasRoomForward {
			break
		}
		if hasRoomBack && hasRoomForward {
			backBuff = (*nums)[backIter] * (*nums)[backIter]
			forwardBuff = (*nums)[forwardIter] * (*nums)[forwardIter]
			if backBuff < forwardBuff {
				*result = append(*result, backBuff)
				backIter -= 1
			} else if backBuff > forwardBuff {
				*result = append(*result, forwardBuff)
				forwardIter += 1
			} else if backBuff == forwardBuff {
				*result = append(*result, backBuff)
				*result = append(*result, forwardBuff)
				forwardIter += 1
				backIter -= 1
			}
		} else if hasRoomBack && !hasRoomForward {
			backBuff = (*nums)[backIter] * (*nums)[backIter]
			*result = append(*result, backBuff)
			backIter -= 1
		} else if !hasRoomBack && hasRoomForward {
			forwardBuff = (*nums)[forwardIter] * (*nums)[forwardIter]
			*result = append(*result, forwardBuff)
			forwardIter += 1
		}

	}
}

func SquareInNormalOrder(nums *[]int, result *[]int) {
	i := 0
	last := len(*nums) - 1
	for {
		if i > last {
			break
		}
		*result = append(*result, (*nums)[i]*(*nums)[i])
		i += 1
	}
}

func SquareInReverseOrder(nums *[]int, result *[]int) {
	i := 0
	last := len(*nums) - 1
	i = last
	for {
		if i < 0 {
			break
		}
		*result = append(*result, (*nums)[i]*(*nums)[i])
		i -= 1
	}
}
