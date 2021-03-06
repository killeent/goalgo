package goalgo

// Queue is a FIFO data structure.
type Queue struct {
	front *queueElement
	back  *queueElement
	count int
}

type queueElement struct {
	next *queueElement
	data interface{}
}

// Push adds an element to the end of the Queue.
func (q *Queue) Push(data interface{}) {
	elt := &queueElement{data: data, next: nil}
	if q.count == 0 {
		q.front = elt
		q.back = elt
	} else if q.count == 1 {
		q.back = elt
		q.front.next = elt
	} else {
		q.back.next = elt
		q.back = elt
	}
	q.count++
}

// Remove removes and returns the first element in the Queue. If
// the Queue is empty, returns nil
func (q *Queue) Remove() interface{} {
	if q.count == 0 {
		return nil
	}
	data := q.front.data
	q.front = q.front.next
	q.count--
	return data
}

// Count returns the number of elements in the Queue.
func (q *Queue) Count() int {
	return q.count
}
