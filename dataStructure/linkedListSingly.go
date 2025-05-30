package dataStructure

import "fmt"

type NodeSingly[T comparable] struct {
    value T
    next  *NodeSingly[T]
}

type LinkedListSingly[T comparable] struct {
    head *NodeSingly[T]
}

func (l *LinkedListSingly[T]) Append(value T) {
    newNodeSingly := &NodeSingly[T]{value: value}
    if l.head == nil {
        l.head = newNodeSingly
        return
    }

    current := l.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNodeSingly
}

func (l *LinkedListSingly[T]) InsertAt(index int, value T) {
    newNodeSingly := &NodeSingly[T]{value: value}

    if index == 0 {
        newNodeSingly.next = l.head
        l.head = newNodeSingly
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

    newNodeSingly.next = current.next
    current.next = newNodeSingly
}

func (l *LinkedListSingly[T]) Print() {
    current := l.head
    for current != nil {
        fmt.Printf("%v -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

func (l *LinkedListSingly[T]) Delete(value T) {
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