package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value int
}
type LinkedListQueue struct {
	head *Node
	len  int
}

func (queue *LinkedListQueue) Size() int {
	return queue.len
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *LinkedListQueue) Enqueue(value int) {
	newNode := &Node{value: value, next: nil}
	if queue.IsEmpty() {
		queue.head = newNode
	} else {
		node := queue.head
		for node.next != nil {
			node = node.next
		}
		node.next = newNode
	}
	queue.len++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Queue is empty"))
	}
	value := queue.head.value
	if queue.Size() == 1 {
		queue.head = nil
	} else {
		queue.head = queue.head.next
	}
	queue.len--
	return value, nil
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Queue is empty"))
	}

	value := queue.head.value

	return value, nil
}

func main() {

}
