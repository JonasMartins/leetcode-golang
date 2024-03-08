package maxfreqelements

import (
	"log"
	"testing"
)

func TestMaxFrequencyElements(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1, 2, 2, 3, 1, 4}
	result := MaxFreqencyElements(input)
	if result < 0 {
		t.Fail()
	}
}
