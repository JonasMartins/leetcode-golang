package taskscheduler

import (
	"log"
	"testing"
)

func TestLastInterval(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	input := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	result := 0
	t.Run("test #1", func(t *testing.T) {
		result = LastInterval(input, 2)
		if result != 8 {
			t.Fail()
		}
	})

	t.Run("test #2", func(t *testing.T) {
		input = []byte{'A', 'C', 'A', 'B', 'D', 'B'}
		result = LastInterval(input, 1)
		if result != 6 {
			t.Fail()
		}
	})

	t.Run("test #3", func(t *testing.T) {
		input = []byte{'A', 'A', 'A', 'B', 'B', 'B'}
		result = LastInterval(input, 3)
		if result != 10 {
			t.Fail()
		}
	})
}
