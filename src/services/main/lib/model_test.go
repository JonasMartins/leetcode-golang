package lib

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

}
