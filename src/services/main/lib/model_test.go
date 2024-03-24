package lib

import (
	"log"
	"testing"
)

func TestModel(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	t.Run("test #1", func(t *testing.T) {})
}
