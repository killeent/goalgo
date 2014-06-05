package goalgo

import (
	"container/heap"
)

type Item struct {
	data     int
	priority int
}

type PriorityQueue []*Item

func (p *PriorityQueue) Len() int {
	return len(p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p *PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(val int) {
	append(*p, val)
	pos := len(*p) - 1
	hole := (pos / 2) - 1
	for p.Less(pos, hole) && hole >= 0 {
		p.Swap(pos, hole)
		hole, pos = (hole/2)-1, hole
	}
}

func (p *PriorityQueue) Pop() int {
	if len(p) == 0 {
		return nil
	}
	val := *p[0]
	*p = *p[1:]
}
