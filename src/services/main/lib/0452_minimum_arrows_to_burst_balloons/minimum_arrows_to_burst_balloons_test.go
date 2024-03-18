package minimumarrowstoburstballoons

import (
	"log"
	"testing"
)

func TestFindMinArrowsShots(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := [][]int{
		{3, 9},
		{7, 12},
		{3, 8},
		{6, 8},
		{9, 10},
		{2, 9},
		{0, 9},
		{3, 9},
		{0, 6},
		{2, 8},
	}
	result := FindMinArrowsShots(input)
	if result == 0 {
		t.Fail()
	}
}
