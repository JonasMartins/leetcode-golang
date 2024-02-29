package substringwithoutrepeat

import "fmt"

func LengthOfLongestSubstring(s string) int {
	switch len(s) {
	case 0:
		return 0
	case 1:
		return 1
	}
	longestSbs := ""
	subsCandidates := map[string]int{}
	str := []rune(s)
	str = append(str, str[len(str)-1])
	sbsPos, i := 0, 0
	ok, hasRepeated, repeatedPos := false, false, 0
	currLongestSubs := ""
	for i < len(str) {
		hasRepeated, repeatedPos = CheckRepeatedCharsFromSubstring(sbsPos, &longestSbs, &str[i])
		if hasRepeated {
			// inserts only the greatest into the map
			// replaces the old one
			_, ok = subsCandidates[currLongestSubs]
			if ok && subsCandidates[currLongestSubs] < len(longestSbs) {
				delete(subsCandidates, currLongestSubs)
				subsCandidates[longestSbs] = len(longestSbs)
				currLongestSubs = longestSbs
			} else {
				if len(currLongestSubs) == 0 {
					subsCandidates[longestSbs] = len(longestSbs)
					currLongestSubs = longestSbs
				}
			}
			longestSbs = ""
			i = repeatedPos + 1
			sbsPos = i
		}
		longestSbs += fmt.Sprintf("%c", str[i])
		i += 1
	}

	return subsCandidates[currLongestSubs]
}

func CheckRepeatedCharsFromSubstring(currStrPos int, sbs *string, r *rune) (bool, int) {
	for i, y := range *sbs {
		if y == *r {
			return true, currStrPos + i
		}
	}
	return false, -1
}
