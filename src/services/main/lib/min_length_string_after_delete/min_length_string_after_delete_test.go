package minlengthstringafterdelete

import (
	"log"
	"testing"
)

func TestMinimunLength(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := "abbbbbbba"
	result := MinimunLength(input)
	if result < 0 {
		t.Fail()
	}
}
