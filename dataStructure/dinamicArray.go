package dataStructure

import (
	"fmt"
	"math"
)

type DinamicArray[T comparable] struct {
	Size     int
	Capacity int
	array    []T
}

func NewDinamicArray[T comparable]() *DinamicArray[T] {
	return &DinamicArray[T]{
		Size:     0,
		Capacity: 5,
		array:    make([]T, 5),
	}
}

func (d *DinamicArray[T]) Add(data T) {
	if d.Size >= d.Capacity {
		d.grow()
	}

	d.array[d.Size] = data
	d.Size++
}

func (d *DinamicArray[T]) Insert(index int, data T) {
	if index > d.Size {
		return
	}
	
	if d.Size >= d.Capacity {
		d.grow()
	}

	for i := index; i < d.Size; i++ {
		d.array[i+1] = d.array[i]
	}

	d.array[index] = data
	d.Size++
}

func (d *DinamicArray[T]) Delete(data T) {
	for i, a := range d.array {
		if a == data {
			for j := i; i < d.Size-1; i++ {
				d.array[j] = d.array[j+1]
			}
			var zero T
			d.array[d.Size-1] = zero
			d.Size--
			if math.Ceil(float64(d.Size+1)/5.0) > math.Ceil(float64(d.Size)/5.0) {
				d.shrink()
			}
			break
		}
	}
}

func (d *DinamicArray[T]) Search(data T) int {
	for i, a := range d.array {
		if a == data {
			return i
		}
	}
	return -1
}

func (d *DinamicArray[T]) IsEmpty() bool {
	return d.Size == 0
}

func (d *DinamicArray[T]) grow() {
	d.Capacity += 5
	newArray := make([]T, d.Capacity)
	copy(newArray, d.array)
	d.array = newArray
}
func (d *DinamicArray[T]) shrink() {
	d.Capacity -= 5
	newArray := make([]T, d.Capacity)
	copy(newArray, d.array)
	d.array = newArray
}

func (d *DinamicArray[T]) ToString() string {
	str := ""
	for i := 0; i < d.Size; i++ {
		str += fmt.Sprintf("%v, ", d.array[i])
	}
	if str != "" {
		str = fmt.Sprintf("[%s]", str[:len(str)-2])
	} else {
		str = "[]"
	}
	return str
}