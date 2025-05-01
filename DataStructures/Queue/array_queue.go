package main

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	front    int
	rear     int
	array    []int
	inserted int
}

func (queue *ArrayQueue) Init(size int) error {
	if size > 0 {
		queue.array = make([]int, size)
		queue.front = -1
		queue.rear = -1
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0")
	}
}

func (queue *ArrayQueue) Size() int {
	return queue.inserted
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *ArrayQueue) Enqueue(value int) error {
	if queue.inserted == len(queue.array) {
		return errors.New("Queue is full")
	}
	if queue.IsEmpty() {
		queue.front = 0
		queue.rear = 0
	} else {
		queue.rear = (queue.rear + 1) % len(queue.array)
	}
	queue.array[queue.rear] = value
	queue.inserted++
	return nil
}

func (queue *ArrayQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Queue is empty"))
	}

	return queue.array[queue.front], nil
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("Queue is empty")
	}
	value := queue.array[queue.front]
	queue.front = (queue.front + 1) % len(queue.array)
	queue.inserted--

	if queue.inserted == 0 {
		queue.front = -1
		queue.rear = -1
	}
	return value, nil
}

// Another method to get the queue size
func (queue *ArrayQueue) length() int {
	if queue.IsEmpty() {
		return 0
	}
	if queue.rear > queue.front {
		return queue.rear - queue.front + 1
	} else {
		return len(queue.array) - queue.front + queue.rear + 1
	}
}
