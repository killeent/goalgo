package goalgo

import (
	"testing"
)

// TestSingleAddRemove tests adding and removing a single int to the Stack.
func TestSingleAddRemove(t *testing.T) {
	st := new(Stack)
	st.Push(0)
	assertEqual(t, 1, st.Count())
	assertEqual(t, 0, extractInt(t, st.Pop()))
	assertEqual(t, 0, st.Count())
}

// TestMultipleAddRemove tests adding and removing a series of ints to the
// Stack.
func TestMultipleAddRemove(t *testing.T) {
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	assertEqual(t, 3, st.Count())
	assertEqual(t, 3, extractInt(t, st.Pop()))
	assertEqual(t, 2, st.Count())
	assertEqual(t, 2, extractInt(t, st.Pop()))
	assertEqual(t, 1, st.Count())
	assertEqual(t, 1, extractInt(t, st.Pop()))
	assertEqual(t, 0, st.Count())
}

// TestInterleavedPushPop tests interleaved addition and removal of elements
// to the Stack.
func TestInterleavedPushPop(t *testing.T) {
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	assertEqual(t, 2, extractInt(t, st.Pop()))
	st.Push(3)
	assertEqual(t, 3, extractInt(t, st.Pop()))
	assertEqual(t, 1, extractInt(t, st.Pop()))
	assertEqual(t, 0, st.Count())
}

// TestEmptyStackNiReturn tests that trying to remove an element from an empty
// Stack returns nil.
func TestEmptyStackNilReturn(t *testing.T) {
	st := new(Stack)
	if st.Pop() != nil {
		t.Error("Invalid return type")
	}
}

func extractInt(t *testing.T, data interface{}) int {
	value, ok := data.(int)
	if !ok {
		t.Error("Invalid return type")
	}
	return value
}

func assertEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}
}
