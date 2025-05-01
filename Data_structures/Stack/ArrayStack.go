package main

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	array    []int
	inserted int
}

func (stack *ArrayStack) Init(size int) error {
	if size > 0 {
		stack.array = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0")
	}
}

func (stack *ArrayStack) DoubleArray() {
	newArray := make([]int, 2*stack.inserted)
	for i := 0; i < stack.inserted; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
}

func (stack *ArrayStack) Size() int {
	return stack.inserted
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *ArrayStack) Push(value int) {
	if stack.Size() == len(stack.array) {
		stack.DoubleArray()
	}
	//Shifting the elements of the array to the right
	for i := stack.inserted; i > 0; i-- {
		stack.array[i] = stack.array[i-1]
	}

	stack.array[0] = value
	stack.inserted++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.Size() <= 0 {
		return -1, errors.New(fmt.Sprintf("Stack is empty"))
	}
	value := stack.array[0]
	for i := 0; i < stack.inserted-1; i++ {
		stack.array[i] = stack.array[i+1]
	}
	stack.inserted--
	return value, nil
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.Size() <= 0 {
		return -1, errors.New(fmt.Sprintf("Stack is empty"))
	}
	value := stack.array[0]
	return value, nil
}

func main() {

}
