package lib

import (
	"log"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	var n int = 1001
	result := IsPalindrome(n)
	if result {
		t.Fail()
	}

}
