package findpivotint

func PivotInteger(n int) int {
	if n <= 1 {
		return 1
	}
	result := -1
	prev, sumPrev, sumFor := uint32(n), uint32(n), uint32(0)
	for prev > 1 {
		prev -= 1
		PrevSum(&prev, &sumPrev)
		SummationFormula(&prev, &sumFor)
		if sumPrev == sumFor {
			result = int(prev)
			break
		}
	}
	return result
}

func SummationFormula(x *uint32, sum *uint32) {
	if *x == 1 {
		*sum = 1
	}
	*sum = *x * (*x + 1) / 2
}

func PrevSum(x *uint32, sum *uint32) {
	*sum += *x
}
