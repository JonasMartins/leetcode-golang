package bagoftokens

import (
	"sort"
)

func BagOfTokens(tokens []int, power int) int {
	size := len(tokens)
	switch size {
	case 0:
		return 0
	case 1:
		if power >= tokens[0] {
			return 1
		} else {
			return 0
		}

	}
	prevResult, result, resultDirect := 0, 0, 0
	sort.Ints(tokens)
	if power < tokens[0] {
		return 0
	} else {
		last := size - 1
		i := 1
		recPower := power
		resultDirect = CalculateDirectely(&tokens, &recPower)
		for {
			power -= tokens[i-1]
			power += tokens[last]
			prevResult = result
			recPower = power
			result = CalculateScoreFromPower(&i, &tokens, &last, &recPower)
			if result <= prevResult {
				result = prevResult
			}
			last -= 1
			i += 1
			if i >= last {
				break
			}
		}
	}
	if resultDirect > result {
		result = resultDirect
	}
	return result
}

func CalculateDirectely(tokens *[]int, power *int) int {
	result := 0
	iter := 0
	for {
		if *power >= (*tokens)[iter] {
			*power -= (*tokens)[iter]
			result += 1
		}
		iter += 1
		if iter == len(*tokens) {
			break
		}
	}
	return result
}

func CalculateScoreFromPower(i *int, tokens *[]int, size *int, power *int) int {
	result := 0
	iter := *i
	for {
		if *power >= (*tokens)[iter] {
			*power -= (*tokens)[iter]
			result += 1
		}
		if iter >= *size-1 {
			break
		}
		iter += 1
	}

	return result
}
