package taskscheduler

func LastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	hash := make(map[byte]uint64)
	ok := false
	for _, x := range tasks {
		_, ok = hash[x]
		if !ok {
			hash[x] = 1
		} else {
			hash[x] = hash[x] + 1
		}
	}
	max := FindMaxFromHash(&hash, tasks[0])
	maxFreqTimeOccupation := int((hash[max] - 1)) * (n + 1)
	amountOfSameMaxAmountFrequency := FindAmountOfBytesWithSameMaxFrequency(&hash, hash[max])
	if maxFreqTimeOccupation+amountOfSameMaxAmountFrequency < len(tasks) {
		return len(tasks)
	}
	return maxFreqTimeOccupation + amountOfSameMaxAmountFrequency
}

func FindAmountOfBytesWithSameMaxFrequency(hash *map[byte]uint64, max uint64) int {
	count := 0
	for _, x := range *hash {
		if x == max {
			count++
		}
	}
	return count
}

func FindMaxFromHash(hash *map[byte]uint64, first byte) byte {
	max := first
	for i, x := range *hash {
		if x > (*hash)[max] {
			max = i
		}
	}
	return max
}
