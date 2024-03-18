package insertinterval

func Insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	startIndex, finishIndex := 0, 0
	shouldMerge, size := false, len(intervals)
	for i, x := range intervals {
		if i < size-1 {
			if newInterval[0] > x[1] && newInterval[1] < intervals[i+1][0] {
				startIndex, finishIndex = i+1, i+1
				shouldMerge = false
				break
			}
			if newInterval[0] > x[1] && newInterval[1] > intervals[i+1][0] {
				startIndex = i + 1
			}
		}
		if x[0] <= newInterval[0] && x[1] >= newInterval[0] {
			startIndex = i
			shouldMerge = true
		}
		if x[0] <= newInterval[1] && x[1] >= newInterval[0] {
			if i < size-1 {
				// will affect next element too, finishIndex must increase
				if newInterval[0] >= intervals[i+1][0] ||
					newInterval[1] >= intervals[i+1][1] ||
					newInterval[1] >= intervals[i+1][0] {
					continue
				}
			}
			finishIndex = i
			shouldMerge = true
			break
		}
		if i > 0 {
			// belongs in the middle
			if newInterval[0] >= intervals[i-1][1] && newInterval[1] <= intervals[i][0] {
				startIndex, finishIndex = i, i
				shouldMerge = true
				break
			}
		}
	}
	result := [][]int{}
	// fits in the begin or end without merge
	if startIndex == 0 && finishIndex == 0 && !shouldMerge {
		// should insert in first or last
		if newInterval[1] < intervals[0][0] {
			result = append(result, newInterval)
			result = append(result, intervals...)
		} else if newInterval[0] > intervals[size-1][0] {
			intervals = append(intervals, newInterval)
			return intervals
		}
	}
	// fits in middle without merge
	if startIndex != 0 && startIndex == finishIndex && !shouldMerge {
		for i, x := range intervals {
			if i == startIndex {
				result = append(result, newInterval)
			}
			result = append(result, x)
		}
		return result
	}
	if !shouldMerge && startIndex != finishIndex {
		intervals = append(intervals, newInterval)
		return intervals
	}

	if shouldMerge {
		newElement := []int{0, 0}
		if intervals[startIndex][0] <= newInterval[0] {
			newElement[0] = intervals[startIndex][0]
		} else {
			newElement[0] = newInterval[0]
		}
		if intervals[finishIndex][1] >= newInterval[1] {
			newElement[1] = intervals[finishIndex][1]
		} else {
			newElement[1] = newInterval[1]
		}
		for i, x := range intervals {
			if i == startIndex {
				result = append(result, newElement)
				continue
			}
			if i > startIndex && i <= finishIndex {
				continue
			}
			result = append(result, x)
		}
	}

	return result
}
