package algorithms

func RecursionBubbleSort[T Ordered](array *[]T, reverse, backward bool) {
	helperOneBubbleSort(array, reverse, backward, 0)
}
func helperOneBubbleSort[T Ordered](array *[]T, reverse, backward bool, index int) {
	if index >= len(*array)-1 {
		return
	}

	if backward {
		helperTwoBubbleSort(array, reverse, backward, len(*array)-1, index)
	} else {
		helperTwoBubbleSort(array, reverse, backward, 0, index)
	}

	helperOneBubbleSort(array, reverse, backward, index+1)
}
func helperTwoBubbleSort[T Ordered](array *[]T, reverse, backward bool, index, max int) {
	if backward {
		if index <= max {
			return
		}
		if reverse {
			if (*array)[index-1] < (*array)[index] {
				(*array)[index-1], (*array)[index] = (*array)[index], (*array)[index-1]
			}
		} else {
			if (*array)[index-1] > (*array)[index] {
				(*array)[index-1], (*array)[index] = (*array)[index], (*array)[index-1]
			}
		}
		helperTwoBubbleSort(array, reverse, backward, index-1, max)
	} else {
		if index >= len(*array)-max-1 {
			return
		}
		if reverse {
			if (*array)[index] < (*array)[index+1] {
				(*array)[index], (*array)[index+1] = (*array)[index+1], (*array)[index]
			}
		} else {
			if (*array)[index] > (*array)[index+1] {
				(*array)[index], (*array)[index+1] = (*array)[index+1], (*array)[index]
			}
		}
		helperTwoBubbleSort(array, reverse, backward, index+1, max)
	}
}

func RecursionSelectionSort[T Ordered](array *[]T, reverse, backward bool) {
	if backward {
		helperOneSelectionSort(array, reverse, backward, len(*array)-1)
	} else {
		helperOneSelectionSort(array, reverse, backward, 0)
	}
}
func helperOneSelectionSort[T Ordered](array *[]T, reverse, backward bool, index int) {
	temp := index
	if backward {
		if index <= 0 {
			return
		}
		helperTwoSelectionSort(array, reverse, backward, index-1, &temp)
		(*array)[index], (*array)[temp] = (*array)[temp], (*array)[index]
		helperOneSelectionSort(array, reverse, backward, index-1)
	} else {
		if index >= len(*array)-1 {
			return
		}
		helperTwoSelectionSort(array, reverse, backward, index+1, &temp)
		(*array)[index], (*array)[temp] = (*array)[temp], (*array)[index]
		helperOneSelectionSort(array, reverse, backward, index+1)
	}
}
func helperTwoSelectionSort[T Ordered](array *[]T, reverse, backward bool, index int, temp *int) {
	if backward {
		if index < 0 {
			return
		}
		if reverse {
			if (*array)[*temp] > (*array)[index] {
				*temp = index
			}
		} else {
			if (*array)[*temp] < (*array)[index] {
				*temp = index
			}
		}
		helperTwoSelectionSort(array, reverse, backward, index-1, temp)
	} else {
		if index >= len(*array) {
			return
		}
		if reverse {
			if (*array)[*temp] < (*array)[index] {
				*temp = index
			}
		} else {
			if (*array)[*temp] > (*array)[index] {
				*temp = index
			}
		}
		helperTwoSelectionSort(array, reverse, backward, index+1, temp)
	}
}

func RecursionInsertionSort[T Ordered](array *[]T, reverse, backward bool) {
	if backward {
		helperOneInsertionSort(array, reverse, backward, len(*array)-2)
	} else {
		helperOneInsertionSort(array, reverse, backward, 1)
	}
}
func helperOneInsertionSort[T Ordered](array *[]T, reverse, backward bool, index int) {
	if backward {
		if index < 0 {
			return
		}
		temp := (*array)[index]

		jndex := index + 1

		helperTwoInsertionSort(array, reverse, backward, &jndex, temp)

		(*array)[jndex-1] = temp

		helperOneInsertionSort(array, reverse, backward, index-1)
	} else {
		if index >= len(*array) {
			return
		}
		temp := (*array)[index]

		jndex := index - 1

		helperTwoInsertionSort(array, reverse, backward, &jndex, temp)

		(*array)[jndex+1] = temp

		helperOneInsertionSort(array, reverse, backward, index+1)
	}
}
func helperTwoInsertionSort[T Ordered](array *[]T, reverse, backward bool, index *int, temp T) {
	if backward {
		if *index >= len(*array) {
			return
		}
		if reverse {
			if (*array)[*index] <= temp {
				return
			}
		} else {
			if (*array)[*index] >= temp {
				return
			}
		}
		(*array)[*index-1] = (*array)[*index]

		*index++
		helperTwoInsertionSort(array, reverse, backward, index, temp)
	} else {
		if *index < 0 {
			return
		}
		if reverse {
			if (*array)[*index] >= temp {
				return
			}
		} else {
			if (*array)[*index] <= temp {
				return
			}
		}
		(*array)[*index+1] = (*array)[*index]
		*index--
		helperTwoInsertionSort(array, reverse, backward, index, temp)
	}
}