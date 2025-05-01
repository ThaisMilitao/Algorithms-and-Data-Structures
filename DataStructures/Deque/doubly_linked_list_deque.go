package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type DoublyLinkedListDeque struct {
	tail   *Node
	head   *Node
	length int
}

func (deque *DoublyLinkedListDeque) Size() int {
	return deque.length
}

func (deque *DoublyLinkedListDeque) IsEmpty() bool {
	return deque.Size() == 0
}

func (deque *DoublyLinkedListDeque) EnqueueFront(value int) {
	newNode := &Node{value: value, next: nil, prev: nil}
	if deque.IsEmpty() {
		deque.head = newNode
		deque.tail = newNode
	}
	newNode.next = deque.head
	deque.head.prev = newNode
	deque.head = newNode

	deque.length++
}

func (deque *DoublyLinkedListDeque) EnqueueRear(value int) {
	newNode := &Node{value: value, next: nil, prev: nil}
	if deque.IsEmpty() {
		deque.head = newNode
		deque.tail = newNode
	}
	newNode.prev = deque.tail
	deque.tail.next = newNode
	deque.tail = newNode

	deque.length++
}

func (deque *DoublyLinkedListDeque) DequeueFront() (int, error) {
	if deque.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.head.value
	if deque.head == deque.tail {
		deque.head = nil
		deque.tail = nil
	} else {
		new_front := deque.head.next
		new_front.prev = nil
		deque.head = new_front
	}
	deque.length--

	return value, nil
}

func (deque *DoublyLinkedListDeque) DequeueRear() (int, error) {
	if deque.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.tail.value
	if deque.head == deque.tail {
		deque.head = nil
		deque.tail = nil
	} else {
		new_tail := deque.tail.prev
		new_tail.next = nil
		deque.tail = new_tail
	}

	deque.length--

	return value, nil
}

func (deque *DoublyLinkedListDeque) Front() (int, error) {
	if deque.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.head.value

	return value, nil
}

func (deque *DoublyLinkedListDeque) Rear() (int, error) {
	if deque.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.tail.value

	return value, nil
}

func main() {
	deque := &DoublyLinkedListDeque{}

}
