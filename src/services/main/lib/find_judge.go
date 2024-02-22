package lib

func FindJudge(n int, trust [][]int) int {
	result := -1
	nonJudgesCandidates := map[int]bool{}
	judgeCandidadesTrustedAmount := map[int]int{}
	for _, x := range trust {
		nonJudgesCandidates[x[0]] = true
	}
	if len(nonJudgesCandidates) == n {
		return result
	}
	for i := 1; i <= n; i += 1 {
		if !nonJudgesCandidates[i] {
			judgeCandidadesTrustedAmount[i] = 0
		}
	}

	for _, x := range trust {
		if judgeCandidadesTrustedAmount[x[1]] >= 0 {
			judgeCandidadesTrustedAmount[x[1]] += 1
		}
	}
	for x := range judgeCandidadesTrustedAmount {
		if judgeCandidadesTrustedAmount[x] == n-1 {
			// one of the judges candidades have everyone
			// else trusting him
			result = x
			break
		}
	}

	return result
}
