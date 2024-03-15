package productofarrayexceptself

import (
	"log"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []int{1, 1}
	result := ProductExceptSelf(input)
	if len(result) == 0 {
		t.Fail()
	}
}
