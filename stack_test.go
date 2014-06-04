package goalgo

import (
	"testing"
)

// TestSingleAddRemove tests adding and removing a single int to the Stack.
func TestStackSingleAddRemove(t *testing.T) {
	st := new(Stack)
	st.Push(0)
	AssertEqual(t, 1, st.Count())
	AssertEqual(t, 0, ExtractInt(t, st.Pop()))
	AssertEqual(t, 0, st.Count())
}

// TestMultipleAddRemove tests adding and removing a series of ints to the
// Stack.
func TestStackMultipleAddRemove(t *testing.T) {
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	AssertEqual(t, 3, st.Count())
	AssertEqual(t, 3, ExtractInt(t, st.Pop()))
	AssertEqual(t, 2, st.Count())
	AssertEqual(t, 2, ExtractInt(t, st.Pop()))
	AssertEqual(t, 1, st.Count())
	AssertEqual(t, 1, ExtractInt(t, st.Pop()))
	AssertEqual(t, 0, st.Count())
}

// TestInterleavedPushPop tests interleaved addition and removal of elements
// to the Stack.
func TestStackInterleavedPushPop(t *testing.T) {
	st := new(Stack)
	st.Push(1)
	st.Push(2)
	AssertEqual(t, 2, ExtractInt(t, st.Pop()))
	st.Push(3)
	AssertEqual(t, 3, ExtractInt(t, st.Pop()))
	AssertEqual(t, 1, ExtractInt(t, st.Pop()))
	AssertEqual(t, 0, st.Count())
}

// TestEmptyStackNiReturn tests that trying to remove an element from an empty
// Stack returns nil.
func TestEmptyStackNilReturn(t *testing.T) {
	st := new(Stack)
	if st.Pop() != nil {
		t.Error("Invalid return type")
	}
}
