package goalgo

import (
	"testing"
)

// TestSingleAddRemove tests addition and removal of a single
// int from the Queue.
func TestQueueSingleAddRemove(t *testing.T) {
	q := new(Queue)
	q.Push(0)
	AssertEqual(t, 1, q.Count())
	AssertEqual(t, 0, ExtractInt(t, q.Remove()))
	AssertEqual(t, 0, q.Count())
}

// TestMultipleAddRemove tests addition and removal of multiple
// ints from the Queue.
func TestQueueMultipleAddRemove(t *testing.T) {
	q := new(Queue)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	AssertEqual(t, 3, q.Count())
	AssertEqual(t, 1, ExtractInt(t, q.Remove()))
	AssertEqual(t, 2, q.Count())
	AssertEqual(t, 2, ExtractInt(t, q.Remove()))
	AssertEqual(t, 1, q.Count())
	AssertEqual(t, 3, ExtractInt(t, q.Remove()))
	AssertEqual(t, 0, q.Count())
}

// TestInterleavedAddRemove tests interleaved addition and removal
// of ints from the Queue.
func TestQueueInterleavedAddRemove(t *testing.T) {
	q := new(Queue)
	q.Push(1)
	q.Push(2)
	AssertEqual(t, 2, q.Count())
	AssertEqual(t, 1, ExtractInt(t, q.Remove()))
	q.Push(3)
	AssertEqual(t, 2, q.Count())
	AssertEqual(t, 2, ExtractInt(t, q.Remove()))
	AssertEqual(t, 3, ExtractInt(t, q.Remove()))
}

// TestEmptyQueueNilRemove tests that removing an item from an empty
// Queue returns nil.
func TestEmptyQueueNilRemove(t *testing.T) {
	q := new(Queue)
	if q.Remove() != nil {
		t.Error("Invalid return value")
	}
}
