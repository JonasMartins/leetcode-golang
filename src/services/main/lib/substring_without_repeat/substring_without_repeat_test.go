package substringwithoutrepeat

import (
	"log"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	result := LengthOfLongestSubstring("anviaj")
	if result < 0 {
		t.Fail()
	}

}
