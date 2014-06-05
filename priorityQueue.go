package goalgo

type Item struct {
	data     interface{}
	priority int
}

type PriorityQueue []*Item

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].priority < (*p)[j].priority
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(data interface{}) {
	it := data.(*Item)
	*p = append(*p, it)
}

func (p *PriorityQueue) Pop() interface{} {
	pq := *p
	it := pq[len(pq)-1]
	*p = pq[:len(pq)-1]
	return it
}
