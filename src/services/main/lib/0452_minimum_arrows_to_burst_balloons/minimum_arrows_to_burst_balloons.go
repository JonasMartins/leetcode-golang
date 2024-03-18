package minimumarrowstoburstballoons

type Interval struct {
	Start int
	End   int
}

func FindMinArrowsShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	j, i := 0, 1
	hash := map[int]*Interval{}
	hash[0] = &Interval{
		Start: points[0][0],
		End:   points[0][1],
	}
	amountMatchesHash := 0
	for ; i < len(points); i += 1 {
		j = i
		for ; j >= 0; j -= 1 {
			if points[i][1] < points[j][0] || points[i][0] > points[j][1] {
				for _, y := range hash {
					if (points[i][0] >= y.Start && points[i][0] <= y.End) ||
						(points[i][1] >= y.Start && points[i][1] <= y.End) {
						amountMatchesHash += 1
						break
					}
				}
				if amountMatchesHash == 0 {
					hash[len(hash)] = &Interval{
						Start: points[i][0],
						End:   points[i][1],
					}
				}
				amountMatchesHash = 0
			}
		}
	}

	return len(hash)
}
