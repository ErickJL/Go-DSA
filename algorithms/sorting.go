package algorithms

func BubbleSort[T Ordered](array *[]T, reverse, backward bool) {
	if backward {
		for i := 0; i < len(*array)-1; i++ {
			for j := len(*array) - 1; j > i; j-- {
				if reverse {
					if (*array)[j-1] < (*array)[j] {
						(*array)[j-1], (*array)[j] = (*array)[j], (*array)[j-1]
					}
				} else {
					if (*array)[j-1] > (*array)[j] {
						(*array)[j-1], (*array)[j] = (*array)[j], (*array)[j-1]
					}
				}
			}
		}
	} else {
		for i := 0; i < len(*array)-1; i++ {
			for j := 0; j < len(*array)-i-1; j++ {
				if reverse {
					if (*array)[j] < (*array)[j+1] {
						(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
					}
				} else {
					if (*array)[j] > (*array)[j+1] {
						(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
					}
				}
			}
		}
	}
}

func SelectionSort[T Ordered](array *[]T, reverse, backward bool) {
	if backward {
		for i := len(*array) - 1; i > 0; i-- {
			temp := i
			for j := i - 1; j >= 0; j-- {
				if reverse {
					if (*array)[temp] > (*array)[j] {
						temp = j
					}
				} else {
					if (*array)[temp] < (*array)[j] {
						temp = j
					}
				}
			}
			(*array)[i], (*array)[temp] = (*array)[temp], (*array)[i]
		}
	} else {
		for i := 0; i < len(*array)-1; i++ {
			temp := i
			for j := i + 1; j < len(*array); j++ {
				if reverse {
					if (*array)[temp] < (*array)[j] {
						temp = j
					}
				} else {
					if (*array)[temp] > (*array)[j] {
						temp = j
					}
				}
			}
			(*array)[i], (*array)[temp] = (*array)[temp], (*array)[i]
		}
	}
}

func InsertionSort[T Ordered](array *[]T, reverse, backward bool) {
	if backward {
		for i := len(*array) - 2; i >= 0; i-- {
			temp := (*array)[i]
			j := i + 1
			if reverse {
				for j < len(*array) && (*array)[j] > temp {
					(*array)[j-1] = (*array)[j]
					j++
				}
				(*array)[j-1] = temp
			} else {
				for j < len(*array) && (*array)[j] < temp {
					(*array)[j-1] = (*array)[j]
					j++
				}
				(*array)[j-1] = temp
			}
		}
	} else {
		for i := 1; i < len(*array); i++ {
			temp := (*array)[i]
			j := i - 1
			if reverse {
				for j >= 0 && (*array)[j] < temp {
					(*array)[j+1] = (*array)[j]
					j--
				}
				(*array)[j+1] = temp
			} else {
				for j >= 0 && (*array)[j] > temp {
					(*array)[j+1] = (*array)[j]
					j--
				}
				(*array)[j+1] = temp
			}
		}
	}
}