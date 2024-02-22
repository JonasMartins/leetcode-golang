package lib

import (
	"log"
	"testing"
)

func TestFindJudge(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	n := 3
	trust := [][]int{
		{1, 3},
		{2, 3},
	}

	// n = 4
	// [[1,3],[1,4],[2,3],[2,4],[4,3]]
	result := FindJudge(n, trust)
	if result < 0 {
		t.Fail()
	}
}
