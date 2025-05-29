package main

import (
	"fmt"

	"github.com/ErickJL/Go-DSA/dataStructure"
)

func main() {
	useStack()
	useQueue()
	usePriorityQueue()
	UseLinkedList()
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
func UseLinkedList() {
	list := dataStructure.LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Print()

	list.Delete(20)
	list.Print()

	list.InsertAt(1, 20)
	list.Print()
}