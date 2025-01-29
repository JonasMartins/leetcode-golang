package romantointeger

func RomanToInt(s string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	value, comb := 0, 0
	var x byte
	var y byte

	for i := 0; i < len(s); i += 1 {
		x = s[i]
		if i < len(s)-1 {
			y = s[i+1]
			comb = analizeDifferentPatter(string([]byte{x, y}))
			if comb > 0 {
				value += comb
				i += 1
				continue
			} else {
				value += values[rune(x)]
			}
		} else {
			value += values[rune(x)]
		}
	}
	return value
}

func analizeDifferentPatter(s string) int {
	switch s {
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	default:
		return 0
	}
}
