package findpivotint

import (
	"log"
	"testing"
)

func TestFindPivotInteger(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := 4
	result := PivotInteger(input)
	if result < -1 {
		t.Fail()
	}
}
