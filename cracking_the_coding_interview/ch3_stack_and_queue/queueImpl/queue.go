package main

type Queue interface {
	Enqueue(element int)
	Dequeue() int
	IsEmpty() bool
}

type IQueue struct {
	list []int
}

func (q *IQueue) Init() {
	q.list = []int{}
}

func (q *IQueue) Enqueue(element int) {
	q.list = append(q.list, element)
}

func (q *IQueue) Dequeue() int {
	if len(q.list) == 0 {
		return -1
	}
	element := q.list[0]
	q.list = q.list[1:]
	return element
}

func (q *IQueue) IsEmpty() bool {
	return len(q.list) == 0
}
