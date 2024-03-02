package maximunoddbinnumber

func MaximumOddBinaryNumber(s string) string {
	result := []rune{}
	sRunes := []rune(s)
	oneAmount, strSize, i := 0, len(sRunes), 0
	for {
		if i >= strSize {
			break
		}
		if sRunes[i] == '1' {
			oneAmount += 1
		}
		i += 1
	}
	i = 0
	for ; i < strSize; i += 1 {
		if i < oneAmount {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}
	result = result[1:]
	result = append(result, '1')
	return string(result)
}
