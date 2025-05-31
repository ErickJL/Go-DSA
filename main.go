package main

import (
	"fmt"
	"time"

	"github.com/ErickJL/Go-DSA/algorithms"
	"github.com/ErickJL/Go-DSA/dataStructure"
)

func main() {
	// Data Structure
	// fmt.Println()
	// fmt.Println("-------------------------------------------> Stack <------------------------------------------")
	// UseStack()
	// fmt.Println()
	// fmt.Println("-------------------------------------------> Queue <------------------------------------------")
	// UseQueue()
	// fmt.Println()
	// fmt.Println("--------------------------------------> Priority Queue <--------------------------------------")
	// UsePriorityQueue()
	// fmt.Println()
	// fmt.Println("------------------------------------> Linked List Singly <------------------------------------")
	// UseLinkedListSingly()
	// fmt.Println()
	// fmt.Println("------------------------------------> Linked List Doubly <------------------------------------")
	// UseLinkedListDoubly()
	// fmt.Println()
	// fmt.Println("--------------------------------------> Dynamic Array <---------------------------------------")
	// UseDynamicArray()

	// Searching
	// arraySearch := make([]int, 10_000_000)
	// for i := 0; i < len(arraySearch); i++ {
	// 	arraySearch[i] = i+1
	// }
	// fmt.Println()
	// fmt.Println("--------------------------------------> Linear Search <---------------------------------------")
	// TryLinearSearch(arraySearch)
	// fmt.Println()
	// fmt.Println("--------------------------------------> Binary Search <---------------------------------------")
	// TryBinarySearch(arraySearch)
	// fmt.Println()
	// fmt.Println("-----------------------------------> Interpolation Search <-----------------------------------")
	// TryInterpolationSearch(arraySearch)

	// Recursion Searching
	// arraySearch := make([]int, 10_000)
	// for i := 0; i < len(arraySearch); i++ {
	// 	arraySearch[i] = i+1
	// }
	// fmt.Println()
	// fmt.Println("----------------------------------> Recursion Linear Search <----------------------------------")
	// TryRecursionLinearSearch(arraySearch)
	// fmt.Println()
	// fmt.Println("----------------------------------> Recursion Binary Search <----------------------------------")
	// TryRecursionBinarySearch(arraySearch)
	// fmt.Println()
	// fmt.Println("-------------------------------> Recursion Interpolation Search <------------------------------")
	// TryRecursionInterpolationSearch(arraySearch)

	// Sorting
	arraySort := make([]int, 1_000)
	for i := 0; i < len(arraySort)/10; i++ {
		arr := []int{2, 7, 5, 3, 8, 1, 10, 9, 6, 4}
		for j := 0; j < 10; j++ {
			arraySort[i*10+j] = arr[j]+i*10
		}
	}
	fmt.Println()
	fmt.Println("----------------------------------------> Bubble Sort <---------------------------------------")
	TryBubbleSort(arraySort)
	fmt.Println()
	fmt.Println("--------------------------------------> Selection Sort <--------------------------------------")
	TrySelectionSort(arraySort)
	fmt.Println()
	fmt.Println("--------------------------------------> Insertion Sort <--------------------------------------")
	TryInsertionSort(arraySort)
	fmt.Println()
	fmt.Println("-----------------------------------> Recursion Bubble Sort <-----------------------------------")
	TryRecursionBubbleSort(arraySort)
	fmt.Println()
	fmt.Println("----------------------------------> Recursion Selection Sort <---------------------------------")
	TryRecursionSelectionSort(arraySort)
	fmt.Println()
	fmt.Println("---------------------------------> Recursion Insertion Sort <---------------------------------")
	TryRecursionInsertionSort(arraySort)
	fmt.Println()
}

// Data Structure
func UseStack() {
	var s dataStructure.Stack[string]

	fmt.Println(s.Empty())

	s.Push("Naruto Shippuden")
	s.Push("One Piece")
	s.Push("Jujutsu Kaisen")
	s.Push("Kuroko no Basket")
	s.Push("One Punch Man")

	fmt.Println(s.Empty())
	fmt.Println(s)

	fmt.Println()
	fmt.Println(s.Pop())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Peek())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Search("One Piece"))
}
func UseQueue() {
	var s dataStructure.Queue[string]

	fmt.Println(s.Empty())

	s.Enqueue("Naruto Shippuden")
	s.Enqueue("One Piece")
	s.Enqueue("Jujutsu Kaisen")
	s.Enqueue("Kuroko no Basket")
	s.Enqueue("One Punch Man")

	fmt.Println(s.Empty())
	fmt.Println(s.Size())
	fmt.Println(s)

	fmt.Println()
	fmt.Println(s.Dequeue())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Peek())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Contains("One Piece"))
}
func UsePriorityQueue() {
	s := dataStructure.NewPriorityQueue[string](true)

	fmt.Println(s.Empty())

	s.Enqueue("Naruto Shippuden")
	s.Enqueue("One Piece")
	s.Enqueue("Jujutsu Kaisen")
	s.Enqueue("Kuroko no Basket")
	s.Enqueue("One Punch Man")

	fmt.Println(s.Empty())
	fmt.Println(s.Size())
	fmt.Println(s)

	fmt.Println()
	fmt.Println(s.Dequeue())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Peek())
	fmt.Println(s)
	
	fmt.Println()
	fmt.Println(s.Contains("One Piece"))
}
func UseLinkedListSingly() {
	list := dataStructure.LinkedListSingly[string]{}

	list.Append("Naruto Shippuden")
	list.Append("One Piece")
	list.Append("Jujutsu Kaisen")
	list.Append("Kuroko no Basket")
	list.Append("One Punch Man")
	list.Print()

	list.Delete("Jujutsu Kaisen")
	list.Print()

	list.InsertAt(2, "Jujutsu Kaisen")
	list.Print()
}
func UseLinkedListDoubly() {
	list := dataStructure.LinkedListDoubly[string]{}

	list.Append("Naruto Shippuden")
	list.Append("One Piece")
	list.Append("Jujutsu Kaisen")
	list.Append("Kuroko no Basket")
	list.Append("One Punch Man")
	list.PrintForward()

	list.Delete("Jujutsu Kaisen")
	list.PrintForward()

	list.InsertAt(2, "Jujutsu Kaisen")
	list.PrintForward()
	list.PrintBackward()
}
func UseDynamicArray() {
	d := dataStructure.NewDinamicArray[string]()

	d.Add("Naruto Shippuden")
	d.Add("One Piece")
	d.Add("Jujutsu Kaisen")
	d.Add("Kuroko no Basket")
	d.Add("One Punch Man")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)
	fmt.Println(d.Search("Kuroko no Basket"))

	d.Delete("Kuroko no Basket")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)
	fmt.Println(d.Search("Kuroko no Basket"))

	d.Insert(3, "Kuroko no Basket")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)

	d.Add("Attack on Titan")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)
	
	d.Add("Bleach")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)

	d.Delete("Bleach")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)

	d.Delete("Attack on Titan")
	fmt.Println(d.ToString())
	fmt.Printf("size: %d\n", d.Size)
	fmt.Printf("capacity: %d\n", d.Capacity)
	fmt.Printf("empty: %t\n", d.IsEmpty())
}

// Search
func TryLinearSearch(array []int) {
	
	start := time.Now()
	index := algorithms.LinearSearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}
func TryBinarySearch(array []int) {
	start := time.Now()
	index := algorithms.BinarySearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}
func TryInterpolationSearch(array []int) {
	start := time.Now()
	index := algorithms.InterpolationSearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}

// Recursion Search
func TryRecursionLinearSearch(array []int) {
	
	start := time.Now()
	index := algorithms.RecursionLinearSearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}
func TryRecursionBinarySearch(array []int) {
	start := time.Now()
	index := algorithms.RecursionBinarySearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}
func TryRecursionInterpolationSearch(array []int) {
	start := time.Now()
	index := algorithms.RecursionInterpolationSearch(array, 7_777_777)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}

// Sort
func TryBubbleSort(array []int) {
	start := time.Now()
	algorithms.BubbleSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}
func TrySelectionSort(array []int) {
	start := time.Now()
	algorithms.SelectionSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}
func TryInsertionSort(array []int) {
	start := time.Now()
	algorithms.InsertionSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}

// Recursion Sort
func TryRecursionBubbleSort(array []int) {
	start := time.Now()
	algorithms.RecursionBubbleSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}
func TryRecursionSelectionSort(array []int) {
	start := time.Now()
	algorithms.RecursionSelectionSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}
func TryRecursionInsertionSort(array []int) {
	start := time.Now()
	algorithms.RecursionInsertionSort(&array, false, false)
	duration := time.Since(start)
	fmt.Println("Duration:", duration.Nanoseconds(), "ns")
	// fmt.Println(array)
}