package main

import (
	"fmt"

	"github.com/ErickJL/Go-DSA/dataStructure"
)

func main() {
	fmt.Println()
	fmt.Println("-------------------------------------------> Stack <------------------------------------------")
	useStack()
	fmt.Println()
	fmt.Println("-------------------------------------------> Queue <------------------------------------------")
	useQueue()
	fmt.Println()
	fmt.Println("--------------------------------------> Priority Queue <--------------------------------------")
	usePriorityQueue()
	fmt.Println()
	fmt.Println("------------------------------------> Linked List Singly <------------------------------------")
	useLinkedListSingly()
	fmt.Println()
	fmt.Println("------------------------------------> Linked List Doubly <------------------------------------")
	useLinkedListDoubly()
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
	s := dataStructure.PriorityQueue[string]{
		Queue: make([]string, 0),
		ReverseOrder: false,
	} 

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