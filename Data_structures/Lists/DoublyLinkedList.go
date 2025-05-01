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

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int //Number of elements inserted into the array
}

func (list *DoublyLinkedList) Size() int {
	return list.length
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
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

func (list *DoublyLinkedList) Set(value int, index int) error {
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

func (list *DoublyLinkedList) Add(value int) {
	newNode := &Node{value: value, next: nil, prev: nil}
	if list.Size() == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.length++
}

func (list *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.length {
		newNode := &Node{value: value, next: nil, prev: nil}
		if list.Size() == 0 || index == list.Size() {
			list.Add(value)
			return nil
		} else {
			if index == 0 {
				list.head.prev = newNode
				newNode.next = list.head
				list.head = newNode
			} else {
				aux := list.head
				for i := 0; i < index-1; i++ {
					aux = aux.next
				}
				newNode.next = aux.next
				newNode.prev = aux
				aux.next.prev = newNode
				aux.next = newNode
			}
			list.length++
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 && index > list.length {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	} else {
		if index == (list.Size() - 1) {
			list.tail = list.tail.prev
			list.tail.next = nil
		} else if index == 0 {
			list.head = list.head.next
			list.head.prev = nil
		} else {
			aux := list.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}
			aux.prev.next = aux.next
			aux.next.prev = aux.prev
		}
		list.length--
		return nil
	}
}

func (list *DoublyLinkedList) Reverse() {
	current := list.head

	for current != nil {
		prev := current.prev
		next := current.next
		current.next = prev
		current.prev = next
		current = next
	}
	new_head := list.tail
	new_tail := list.head
	list.head = new_head
	list.tail = new_tail

}
