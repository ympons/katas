package main

type Queue interface {
	Enqueue(int)
	Dequeue() (int, bool)
	IsEmpty() bool
}

type qnode struct {
	value int
	next *qnode
}

type queue_ struct {
	head *qnode
	tail *qnode
	n int
}

func NewQueue() *queue_{
	return &queue_{}
}

func (q *queue_) Enqueue(v int) {
	node := &qnode{value: v}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.n++
}

func (q *queue_) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	node := q.head
	q.head = node.next
	if q.head == nil {
		q.tail = nil
	}
	q.n--

	return node.value, true
}

func (q queue_) IsEmpty() bool {
	return q.n == 0
}
