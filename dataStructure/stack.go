package dataStructure

type Stack[T comparable] []T

func (s *Stack[T]) Push(data T) {
	*s = append(*s, data)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(*s) - 1
	result := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return result, true
}
func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}

	result := (*s)[len(*s)-1]
	return result, true
}
func (s *Stack[T]) Search(data T) int {
	for i, s := range *s {
		if s == data {
			return i
		}
	}
	return -1
}
func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}