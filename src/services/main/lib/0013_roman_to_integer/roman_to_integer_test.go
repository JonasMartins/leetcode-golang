package romantointeger

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()
	inputAndResults := map[string]int{
		"I":           1,
		"II":          2,
		"CD":          400,
		"V":           5,
		"MM":          2000,
		"IV":          4,
		"III":         3,
		"LIV":         54,
		"LVIII":       58,
		"MCMXCIV":     1994,
		"MMXXIII":     2023,
		"MMMCMXCIX":   3999,
		"MMDCCCLXXIX": 2879,
		"MCMXCIX":     1999,
		"MCDLXXVI":    1476,
	}
	value := 0
	for key, val := range inputAndResults {
		value = RomanToInt(key)
		assert.Equal(t, val, value)
	}
}
