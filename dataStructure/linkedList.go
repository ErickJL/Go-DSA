package dataStructure

import "fmt"

type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
}

func (l *LinkedList) Append(value int) {
    newNode := &Node{value: value}
    if l.head == nil {
        l.head = newNode
        return
    }

    current := l.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

func (l *LinkedList) InsertAt(index, value int) {
    newNode := &Node{value: value}

    if index == 0 {
        newNode.next = l.head
        l.head = newNode
        return
    }

    current := l.head
    for i := 0; current != nil && i < index-1; i++ {
        current = current.next
    }

    if current == nil {
        fmt.Println("Index out of range")
        return
    }

    newNode.next = current.next
    current.next = newNode
}

func (l *LinkedList) Print() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()
}

func (l *LinkedList) Delete(value int) {
    if l.head == nil {
        return
    }

    if l.head.value == value {
        l.head = l.head.next
        return
    }

    current := l.head
    for current.next != nil && current.next.value != value {
        current = current.next
    }

    if current.next != nil {
        current.next = current.next.next
    }
}