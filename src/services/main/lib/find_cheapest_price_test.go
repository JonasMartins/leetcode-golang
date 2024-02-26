package lib

import (
	"log"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	n := 6
	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 4, 400},
		{2, 0, 100},
		{4, 1, 200},
		{1, 3, 600},
		{2, 3, 200},
		{0, 5, 500},
		{5, 3, 10},
	}
	src := 0
	dst := 3
	k := 1
	result := FindCheapestPrice(n, flights, src, dst, k)
	if result < 0 {
		t.Fail()
	}
}
