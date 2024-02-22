package lib

import (
	"fmt"
	"math"
)

func RangeBitwiseAnd(left int, right int) int {
	if left == 0 || left > right {
		return 0
	}
	if left == right {
		return left
	}

	leftStr := fill32Pos(fmt.Sprintf("%b", left))
	rightStr := fill32Pos(fmt.Sprintf("%b", right))
	j, pos := 0, 0
	buff := ""
	for ; j < len(leftStr)-1; j += 1 {
		buff = string(leftStr[j])
		if buff == "1" {
			pos = j
			break
		}
	}
	j = 0
	for ; j < pos; j += 1 {
		buff = string(rightStr[j])
		if buff == "1" {
			pos = 0
			break
		}
	}

	if leftStr[j] != rightStr[j] || pos == 0 {
		return 0
	}
	j = pos + 1
	stillMatch := true
	prefix := "1"
	for ; j < len(leftStr); j += 1 {
		stillMatch = leftStr[j] == rightStr[j]
		if stillMatch {
			prefix += string(leftStr[j])
		} else {
			break
		}
	}
	for ; j < len(leftStr); j += 1 {
		prefix += "0"
	}

	return convertBinStrToInt(prefix)
}

func fill32Pos(n string) string {
	i := 0
	buff := ""
	for ; i < (32 - len(n)); i += 1 {
		buff += "0"
	}
	return buff + n
}

func convertBinStrToInt(str string) int {
	sum := 0
	i, j := len(str)-1, 1

	for ; i >= 0; i -= 1 {
		if string(str[i]) == "1" {
			sum += int(math.Pow(2.0, float64(j-1)))
		}
		j += 1
	}
	return sum
}
