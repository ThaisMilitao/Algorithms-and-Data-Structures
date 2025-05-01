package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type LinkedListStack struct {
	head   *Node
	length int
}

func (stack *LinkedListStack) Size() int {
	return stack.length
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.Size() <= 0
}

func (stack *LinkedListStack) Push(value int) {
	new_node := &Node{value: value, next: nil}

	if stack.Size() > 0 {
		next_element := stack.head
		new_node.next = next_element
	}

	stack.head = new_node
	stack.length++
}

func (stack *LinkedListStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Stack is empty"))
	}
	return stack.head.value, nil
}

func (stack *LinkedListStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New(fmt.Sprintf("Stack is empty"))
	}
	value := stack.head.value
	next_node := stack.head.next
	stack.head = next_node
	stack.length--
	return value, nil
}

func balparenteses(par string) bool {
	stack := &LinkedListStack{}
	for _, i := range par {
		if i == '(' {
			stack.Push(0)
		} else if i == ')' {
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func main() {

}
