package regularexpressionmatching

import (
	"log"
	"testing"
)

func TestIsMatching(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	s := "c"
	p := "wwwb*.*...yyyyc*..."
	result := IsMatch(s, p)
	if result {
		t.Fail()
	}
}
