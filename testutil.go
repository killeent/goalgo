package goalgo

import (
	"testing"
)

// ExtractInt attempts to extract an int from an empty interface. If the
// underlying type of the interface variable is not an int, this function
// raises an error on the passed testing variable.
func ExtractInt(t *testing.T, data interface{}) int {
	value, ok := data.(int)
	if !ok {
		t.Error("Invalid return type")
	}
	return value
}

// AssertEqual checks if two ints are equal. If not, raises an error on
// the passed testing variable.
func AssertEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}
}
