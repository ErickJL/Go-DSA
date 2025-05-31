package algorithms

func RecursionLinearSearch[T Ordered](array []T, value T) int {
	return helperLinearSearch(array, value, 0)
}
func helperLinearSearch[T Ordered](array []T, value T, index int) int {
	if index >= len(array) {
		return -1
	}
	if array[index] == value {
		return index
	}
	return helperLinearSearch(array, value, index+1)
}
func RecursionBinarySearch[T Ordered](array []T, value T) int {
	return helperBinarySearch(array, value, 0, len(array)-1)
}
func helperBinarySearch[T Ordered](array []T, value T, low, high int) int {
	if !(low <= high && value <= array[high] && value >= array[low]) {
		return -1
	}

	middle := low + ((high - low) / 2)
	if array[middle] < value {
		return helperBinarySearch(array, value, middle+1, high)
	} else if array[middle] > value {
		return helperBinarySearch(array, value, low, middle-1)
	} else {
		return middle
	}
}
func RecursionInterpolationSearch(array []int, value int) int {
	return helperInterpolationSearch(array, value, 0, len(array)-1)
}
func helperInterpolationSearch(array []int, value int, low, high int) int {
	if !(low <= high && value <= array[high] && value >= array[low]) {
		return -1
	}

	probe := low + (value-array[low])*((high-low)/(array[high]-array[low]))
	if array[probe] < value {
		return helperInterpolationSearch(array, value, probe+1, high)
	} else if array[probe] > value {
		return helperInterpolationSearch(array, value, low, probe-1)
	} else {
		return probe
	}
}