package goalgo

import (
	"testing"
)

// TestSingleAddRemove tests adding and removing a single int to the Stack.
func TestSingleAddRemove(t *testing.T) {
	st := new(Stack)
	st.Push(0)
	if st.Count() != 1 {
		t.Errorf("Expected count: %d, Actual: %d", 1, st.Count())
	}
	data, ok := st.Pop().(int)
	if !ok {
		t.Errorf("Invalid return type")
	}
	if data != 0 {
		t.Errorf("Expected value: %d, Actual: %d", 0, data)
	}
	if st.Count() != 0 {
		t.Errorf("Expected count: %d, Actual: %d", 0, st.Count())
	}
}
