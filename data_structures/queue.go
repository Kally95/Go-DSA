package data_structures

type Queue[T any] struct {
	Nodes  []T
	Length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	q.Nodes = append(q.Nodes, v)
	q.Length++
}

func (q *Queue[T]) QueueSize() int {
	return q.Length
}

func (q *Queue[T]) Dequeue() (t T) {
	if q.Length == 0 {
		return t
	}
	toRemove := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return toRemove
}
