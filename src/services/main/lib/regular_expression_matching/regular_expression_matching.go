package regularexpressionmatching

func IsMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	sRune := []rune(s)
	pRune := []rune(p)
	it := uint32(0)
	rules := GetRulesFromPattern(&pRune)
	if len(*rules) == 1 && (*rules)[0][len((*rules)[0])-1] != '*' {
		return CompareTwoStrings(&sRune, &pRune, &it)
	}

	// given all the rules, everyone of them should match the string in order
	// if so, will return true, false otherwise

	return false
}

// Given the pattern, returns each rule in order
func GetRulesFromPattern(p *[]rune) *[]string {
	rules := []string{}
	aux, buff := "", ""
	i := 0
	for ; i < len(*p); i += 1 {
		buff += string((*p)[i])
		if (*p)[i] == '*' {
			aux = buff[len(buff)-2:]
			buff = buff[:len(buff)-2]
			if len(buff) > 0 {
				rules = append(rules, buff)
			}
			rules = append(rules, aux)
			buff = ""
		}
	}
	if len(buff) > 0 {
		rules = append(rules, buff)
	}

	return &rules
}

func CompareTwoStrings(str1 *[]rune, str2 *[]rune, it *uint32) bool {
	size := uint32(len(*str1))
	if int(size) != len(*str2) {
		return false
	}
	*it = 0
	for {
		if (*str1)[*it] != (*str2)[*it] && (*str1)[*it] != '.' {
			return false
		}
		if *it == size-1 {
			break
		}
	}
	return true
}
