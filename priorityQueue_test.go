package goalgo

import (
	"container/heap"
	"testing"
)

// TestHeapInit tests that creating a PriorityQueue from an existing
// collection of Items puts them in the proper order.
func TestPQInit() {
	items := []*Item{&Item{data: 1, priority: 1},
		&Item{data: 2, priority: 2},
		&Item{data: 3, priority: 3}}
	heap.Init(items)
	AssertEqual(t, expected, actual)
}

func ExtractValueFromItem(it *Item) {

}
