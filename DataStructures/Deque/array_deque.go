package main

import (
	"errors"
	"fmt"
)

type ArrayDeque struct {
	array    []int
	inserted int
	front    int
	rear     int
}

func (deque *ArrayDeque) Init(size int) error {
	if size > 0 {
		deque.array = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0")
	}
}

func (deque *ArrayDeque) DoubleArray() {
	newArray := make([]int, 2*len(deque.array))

	for i := 0; i < deque.inserted; i++ {
		newArray[i] = deque.array[(deque.front+i)%len(deque.array)]
	}

	deque.array = newArray
	deque.front = 0
	deque.rear = deque.inserted
}

func (deque *ArrayDeque) Size() int {
	return deque.inserted
}

func (deque *ArrayDeque) IsEmpty() bool {
	return deque.Size() == 0
}

func (deque *ArrayDeque) EnqueueFront(value int) {
	if deque.inserted == len(deque.array) {
		deque.DoubleArray()
	}

	deque.front = (deque.front - 1 + len(deque.array)) % len(deque.array)
	deque.array[deque.front] = value
	deque.inserted++
}

func (deque *ArrayDeque) EnqueueRear(value int) {
	if deque.inserted == len(deque.array) {
		deque.DoubleArray()
	}

	deque.rear = (deque.rear + 1) % len(deque.array)
	deque.array[deque.rear] = value
	deque.inserted++
}

func (deque *ArrayDeque) DequeueFront() (int, error) {
	if deque.Size() == 0 {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.array[deque.front]
	deque.front = (deque.front + 1) % len(deque.array)
	deque.inserted--

	if deque.inserted == 0 {
		deque.front = -1
		deque.rear = -1
	}

	return value, nil
}
func (deque *ArrayDeque) DequeueRear() (int, error) {
	if deque.Size() == 0 {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.array[deque.rear]
	deque.rear = (deque.rear - 1 + len(deque.array)) % len(deque.array)

	deque.inserted--

	if deque.inserted == 0 {
		deque.front = -1
		deque.rear = -1
	}

	return value, nil
}
func (deque *ArrayDeque) Front() (int, error) {
	if deque.Size() == 0 {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.array[deque.front]
	return value, nil
}
func (deque *ArrayDeque) Rear() (int, error) {
	if deque.Size() == 0 {
		return -1, errors.New(fmt.Sprintf("Deque is empty"))
	}
	value := deque.array[deque.rear]
	return value, nil
}

func main() {

}
