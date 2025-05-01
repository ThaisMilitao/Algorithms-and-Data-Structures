package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head   *Node
	length int //Number of elements inserted into the array
}

func (list *LinkedList) Size() int {
	return list.length
}

func (list *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.length {
		current_node := list.head
		for i := 0; i < index; i++ {
			current_node = current_node.next
		}
		return current_node.value, nil
	} else {
		return -1, errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *LinkedList) Set(value int, index int) error {
	if index >= 0 && index < list.length {
		current_node := list.head
		for i := 0; i < index; i++ {
			current_node = current_node.next
		}
		current_node.value = value
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *LinkedList) Add(value int) {
	newNode := &Node{value: value, next: nil}
	if list.Size() == 0 {
		list.head = newNode
	} else {
		current_node := list.head
		for current_node.next != nil {
			current_node = current_node.next
		}
		current_node.next = newNode
	}
	list.length++
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.length {
		newNode := &Node{value: value, next: nil}
		if index == 0 {
			newNode.next = list.head
			list.head = newNode
		} else {
			current_node := list.head
			for i := 0; i < index-1; i++ {
				current_node = current_node.next
			}
			newNode.next = current_node.next
			current_node.next = newNode
		}
		list.length++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index <= list.length {
		if index == 0 {
			list.head = list.head.next
		} else {
			current_node := list.head
			for i := 0; i < index-1; i++ {
				current_node = current_node.next
			}
			current_node.next = current_node.next.next
		}
		list.length--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *LinkedList) Reverse() {
	var prev *Node
	current := list.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}
