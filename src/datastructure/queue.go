package datastructure

type Queue struct {
	start *LinkedNode
	end   *LinkedNode
}

func (q *Queue) Add(value int) {
	node := LinkedNode{value: value}

	if q.end != nil {
		q.end.next = &node
	}

	q.end = &node

	if q.start == nil {
		q.start = &node
		q.end = q.start
	}
}

func (q *Queue) Remove() *LinkedNode {
	if q.start == nil {
		return nil
	}
	node := q.start
	q.start = q.start.next

	if q.start == nil {
		q.end = nil
	}

	node.next = nil
	return node
}

func (q *Queue) IsEmpty() bool {
	return q.start == nil
}
