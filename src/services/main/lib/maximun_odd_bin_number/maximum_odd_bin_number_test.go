package maximunoddbinnumber

import (
	"log"
	"testing"
)

func TestMaximumOddBinaryNumber(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := "000100000010000000"
	result := MaximumOddBinaryNumber(input)
	if len(result) < 1 {
		t.Fail()
	}

}
