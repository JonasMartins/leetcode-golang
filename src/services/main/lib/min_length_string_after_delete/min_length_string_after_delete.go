package minlengthstringafterdelete

func MinimunLength(s string) int {
	buff := []rune(s)
	for len(buff) > 0 {
		if !CheckTrimPrefixSufixAvailability(&buff) {
			break
		} else {
			if !CheckIfThereIsDifferentChars(&buff) {
				buff = buff[:0]
			} else {
				TrimSufixPrefix(&buff)
			}
		}
	}
	return len(buff)
}

func CheckIfThereIsDifferentChars(s *[]rune) bool {
	fst := (*s)[0]
	for i := 1; i < len((*s)); i += 1 {
		if (*s)[i] != fst {
			return true
		}
	}
	return false
}

func CheckTrimPrefixSufixAvailability(s *[]rune) bool {
	if len(*s) == 1 {
		return false
	}
	return (*s)[0] == (*s)[len(*s)-1]
}

func TrimSufixPrefix(s *[]rune) {
	currRune := (*s)[0]
	startTrim := 0
	finishTrim := len(*s) - 1
	startIsEqual := (*s)[startTrim] == currRune
	finishIsEqual := (*s)[finishTrim] == currRune
	for finishTrim >= 0 && startTrim < len(*s) {
		if startIsEqual {
			startTrim += 1
		}
		if finishIsEqual {
			finishTrim -= 1
		}
		if startTrim >= finishTrim || (!startIsEqual && !finishIsEqual) {
			break
		}
		startIsEqual = (*s)[startTrim] == currRune
		finishIsEqual = (*s)[finishTrim] == currRune
	}
	*s = (*s)[:finishTrim+1]
	if len(*s) >= startTrim {
		*s = (*s)[startTrim:]
	}
}
