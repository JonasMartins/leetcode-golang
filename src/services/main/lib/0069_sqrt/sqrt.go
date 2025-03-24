package sqrt

func MySqrt(x int) int {
	l, r, m := 0, x, 0
	res := 0
	for l <= r {
		m = l + ((r - l) / 2)
		if m*m > x {
			r = m - 1
		} else if m*m < x {
			l = m + 1
			res = m
		} else {
			return m
		}
	}
	return res
}
