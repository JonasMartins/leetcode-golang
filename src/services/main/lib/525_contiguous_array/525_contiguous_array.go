package contiguousarray

type Info struct {
	firstIndxZero int
	firstIndexOne int
	sum           int
	result        int
}

func FindMaxLength(nums []int) int {
	data := FindMaxSubArr(&nums, 0)
	if data.result > 0 {
		return data.result
	}
	var data2 *Info = nil
	if data.sum > 0 {
		data2 = FindMaxSubArr(&nums, data.firstIndxZero)
	} else {
		data2 = FindMaxSubArr(&nums, data.firstIndexOne)
	}
	return data2.result
}

func FindMaxSubArr(nums *[]int, start int) *Info {
	firstOne, firstZero, sum, result := -1, -1, 0, 0
	i := start
	for ; i < len(*nums); i += 1 {
		if (*nums)[i] == 0 {
			firstZero = i
		} else {
			firstOne = i
		}
		if firstOne != -1 && firstZero != -1 {
			break
		}
	}
	i = start
	for ; i < len(*nums); i += 1 {
		if (*nums)[i] == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		if sum == 0 {
			result = i
		}
	}
	r := Info{
		sum:           sum,
		firstIndxZero: firstZero,
		firstIndexOne: firstOne,
		result:        result,
	}
	if r.result > 0 {
		r.result += 1
		r.result = r.result - start
	}
	return &r
}

func HandleOnesDiff(nums *[]int) int {
	size := len(*nums)
	seq1, seq0, maxSeq := 0, 0, 0
	i, firstZeroIndex := 0, 0
	for ; i < size; i += 1 {
		if (*nums)[i] == 0 {
			firstZeroIndex = i
			break
		}
	}
	i = firstZeroIndex
	for ; i < size; i += 1 {
		if (*nums)[i] == 0 {
			seq0 += 1
		} else {
			seq1 += 1
		}
		if seq1 > seq0 || i == size-1 {
			if maxSeq < seq0 {
				maxSeq = seq0
			}
			seq1, seq0 = 0, 0
		}
	}

	return maxSeq * 2
}

func HandleZerosDiff(nums *[]int) int {
	size := len(*nums)
	seq1, seq0, maxSeq := 0, 0, 0
	i, firstOneIndex := 0, 0
	for ; i < size; i += 1 {
		if (*nums)[i] == 1 {
			firstOneIndex = i
			break
		}
	}
	i = firstOneIndex
	for ; i < size; i += 1 {
		if (*nums)[i] == 0 {
			seq0 += 1
		} else {
			seq1 += 1
		}
		if seq0 > seq1 || i == size-1 {
			if maxSeq < seq1 {
				maxSeq = seq1
			}
			seq1, seq0 = 0, 0
		}
	}
	return maxSeq * 2
}
