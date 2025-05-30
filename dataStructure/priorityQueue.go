package dataStructure

type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

type PriorityQueue[T Ordered] struct {
	Queue        []T
	ReverseOrder bool
}

func (q *PriorityQueue[T]) Enqueue(data T) {
	(*q).Queue = append((*q).Queue, data)
	q.sort()
}

func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	if len((*q).Queue) == 0 {
		var zero T
		return zero, false
	}

	result := ((*q).Queue)[0]
	(*q).Queue = ((*q).Queue)[1:]
	return result, true
}
func (q *PriorityQueue[T]) Peek() (T, bool) {
	if len((*q).Queue) == 0 {
		var zero T
		return zero, false
	}

	result := ((*q).Queue)[0]
	return result, true
}
func (q *PriorityQueue[T]) Contains(data T) bool {
	for _, s := range (*q).Queue {
		if s == data {
			return true
		}
	}
	return false
}
func (q *PriorityQueue[T]) Empty() bool {
	return len((*q).Queue) == 0
}

func (q *PriorityQueue[T]) Size() int {
	return len((*q).Queue)
}

func (q *PriorityQueue[T]) sort() {
	for i := 0; i < len((*q).Queue); i++ {
		for j := i + 1; j < len((*q).Queue); j++ {
			if (*q).ReverseOrder {
				if ((*q).Queue)[j] > ((*q).Queue)[i] {
					((*q).Queue)[i], ((*q).Queue)[j] = ((*q).Queue)[j], ((*q).Queue)[i]
				}
			} else {
				if ((*q).Queue)[j] < ((*q).Queue)[i] {
					((*q).Queue)[i], ((*q).Queue)[j] = ((*q).Queue)[j], ((*q).Queue)[i]
				}
			}
		}
	}
}