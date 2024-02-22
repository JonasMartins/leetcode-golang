package lib

import (
	"log"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	result := RangeBitwiseAnd(2, 2)
	if result < 0 {
		t.Fail()
	}

}
