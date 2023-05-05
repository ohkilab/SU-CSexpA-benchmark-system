package worker

import "sync"

type Queue[T any] struct {
	list []*T
	sync.Mutex
}

func (q *Queue[T]) Push(t *T) {
	q.Lock()
	defer q.Unlock()
	q.list = append(q.list, t)
}

func (q *Queue[T]) Pop() *T {
	q.Lock()
	defer q.Unlock()
	if len(q.list) == 0 {
		return nil
	}
	head := q.list[0]
	q.list = q.list[1:]
	return head
}
