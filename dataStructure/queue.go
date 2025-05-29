package dataStructure

type Queue[T comparable] []T

func (q *Queue[T]) Enqueue(data T) {
	*q = append(*q, data)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var zero T
		return zero, false
	}

	result := (*q)[0]
	*q = (*q)[1:]
	return result, true
}
func (q *Queue[T]) Peek() (T, bool) {
	if len(*q) == 0 {
		var zero T
		return zero, false
	}

	result := (*q)[len(*q)-1]
	return result, true
}
func (q *Queue[T]) Contains(data T) bool {
	for _, s := range *q {
		if s == data {
			return true
		}
	}
	return false
}
func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}