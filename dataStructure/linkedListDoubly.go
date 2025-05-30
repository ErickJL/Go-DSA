package dataStructure

import "fmt"

type NodeDoubly[T comparable] struct {
    value T
    next  *NodeDoubly[T]
    prev  *NodeDoubly[T]
}

type LinkedListDoubly[T comparable] struct {
    head *NodeDoubly[T]
    tail *NodeDoubly[T]
}

func (l *LinkedListDoubly[T]) Append(value T) {
    newNodeDoubly := &NodeDoubly[T]{value: value}
    if l.head == nil {
        l.head = newNodeDoubly
        l.tail = newNodeDoubly
        return
    }

    l.tail.next = newNodeDoubly
    newNodeDoubly.prev = l.tail
    l.tail = newNodeDoubly
}

func (l *LinkedListDoubly[T]) InsertAt(index int, value T) {
    newNodeDoubly := &NodeDoubly[T]{value: value}

    if index == 0 {
        newNodeDoubly.next = l.head
        if l.head != nil {
            l.head.prev = newNodeDoubly
        }
        l.head = newNodeDoubly
        if l.tail == nil {
            l.tail = newNodeDoubly
        }
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

    newNodeDoubly.next = current.next
    newNodeDoubly.prev = current
    if current.next != nil {
        current.next.prev = newNodeDoubly
    } else {
        l.tail = newNodeDoubly
    }
    current.next = newNodeDoubly
}

func (l *LinkedListDoubly[T]) PrintForward() {
    current := l.head
    for current != nil {
        fmt.Printf("%v -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

func (l *LinkedListDoubly[T]) PrintBackward() {
    current := l.tail
    for current != nil {
        fmt.Printf("%v <- ", current.value)
        current = current.prev
    }
    fmt.Println("nil")
}

func (l *LinkedListDoubly[T]) Delete(value T) {
    current := l.head
    for current != nil && current.value != value {
        current = current.next
    }

    if current == nil {
        return
    }

    if current.prev != nil {
        current.prev.next = current.next
    } else {
        l.head = current.next
    }

    if current.next != nil {
        current.next.prev = current.prev
    } else {
        l.tail = current.prev
    }
}