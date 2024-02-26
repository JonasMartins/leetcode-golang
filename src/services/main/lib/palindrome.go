package lib

import (
	"math/big"
)

func IsPalindrome(x int) bool {

	var n big.Int
	n.SetInt64(int64(x))
	singleDigit := n.Cmp(big.NewInt(0))
	switch singleDigit {
	case 0:
		return true
	case -1:
		return false
	}
	str := n.String()
	if len(str) == 2 {
		return str[0] == str[1]
	}
	i := 0
	j := len(str) - 1
	isPal := true
	for {
		if i >= j {
			if str[i] != str[j] {
				isPal = false
			}
			break
		}
		if str[i] != str[j] {
			isPal = false
			break
		}
		i += 1
		j -= 1
	}

	return isPal
}
