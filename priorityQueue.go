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
	*p = append(*p, val)
	pos := len(p) - 1
	hole := (pos / 2) - ((pos + 1) % 2)
	for p.Less(pos, hole) && hole >= 0 {
		p.Swap(pos, hole)
		hole, pos = (hole/2)-((hole+1)%2), hole
	}
}

func (p *PriorityQueue) Pop() int {
	if len(p) == 0 {
		return nil
	}
	val := p[0]
	p.Swap(0, len(p)-1)
	*p = p[:len(p)-1]
	pos := 0
	hole := 1
	for hole < len(p) {
		if hole+1 < len(p) && p.Less(hole+1, hole) {
			hole++
		}
		if p.Less(hole, pos) {
			p.Swap(pos, hole)
			hole, pos = (hole*2)+1, hole
		} else {
			break
		}
	}
	return val
}
