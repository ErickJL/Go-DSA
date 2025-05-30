package main

import (
	"fmt"

	"github.com/ErickJL/Go-DSA/dataStructure"
)

func main() {
	// fmt.Println()
	// fmt.Println("-------------------------------------------> Stack <------------------------------------------")
	// useStack()
	// fmt.Println()
	// fmt.Println("-------------------------------------------> Queue <------------------------------------------")
	// useQueue()
	// fmt.Println()
	// fmt.Println("--------------------------------------> Priority Queue <--------------------------------------")
	// usePriorityQueue()
	// fmt.Println()
	// fmt.Println("------------------------------------> Linked List Singly <------------------------------------")
	// useLinkedListSingly()
	// fmt.Println()
	// fmt.Println("------------------------------------> Linked List Doubly <------------------------------------")
	// useLinkedListDoubly()
	fmt.Println()
	fmt.Println("----------------------------------------> Dynamic Array <----------------------------------------")
	useDynamicArray()
	fmt.Println()
}

func useStack() {
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

func useQueue() {
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

func usePriorityQueue() {
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
func useLinkedListSingly() {
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
func useLinkedListDoubly() {
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

func useDynamicArray() {
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