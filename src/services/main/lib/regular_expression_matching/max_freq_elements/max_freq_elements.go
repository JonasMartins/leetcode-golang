package maxfreqelements

func MaxFreqencyElements(nums []int) int {
	result := 0
	if len(nums) == 1 {
		return 1
	}
	hash := map[uint8]uint8{}
	max := uint8(1)
	ux := uint8(1)
	for _, x := range nums {
		ux = uint8(x)
		_, ok := hash[ux]
		if ok {
			hash[ux] += 1
			if hash[ux] > max {
				max = hash[ux]
			}
		} else {
			hash[ux] = uint8(1)
		}
	}
	for _, x := range hash {
		if x == max {
			result += int(max)
		}
	}

	return result
}
