package algorithms

func LinearSearch[T Ordered](array []T, value T) int {
	for i, a := range array {
		if a == value {
			return i
		}
	}
	return -1
}

func BinarySearch[T Ordered](array []T, value T) int {
	low := 0
	high := len(array) - 1

	for low <= high && value <= array[high] && value >= array[low] {
		middle := low + ((high - low) / 2)
		if array[middle] < value {
			low = middle + 1
		} else if array[middle] > value {
			high = middle - 1
		} else {
			return middle
		}
	}

	return -1
}

func InterpolationSearch(array []int, value int) int {
	low := 0
	high := len(array) - 1

	for low <= high && value <= array[high] && value >= array[low] {
		probe := low + (value-array[low])*((high-low)/(array[high]-array[low]))
		if array[probe] < value {
			low = probe + 1
		} else if array[probe] > value {
			high = probe - 1
		} else {
			return probe
		}
	}

	return -1
}