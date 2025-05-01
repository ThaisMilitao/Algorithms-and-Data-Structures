package main

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	array    []int
	inserted int //Number of elements inserted into the array
}

func (list *ArrayList) Init(size int) error {
	if size > 0 {
		list.array = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0")
	}
}

func (list *ArrayList) doubleArray() {
	newArray := make([]int, 2*list.inserted)
	for i := 0; i < list.inserted; i++ {
		newArray[i] = list.array[i]
	}
	list.array = newArray
}

func (list *ArrayList) size() int {
	return list.inserted
}

func (list *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.array[index], nil
	} else {
		return -1, errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *ArrayList) Set(value int, index int) error {
	if index >= 0 && index < list.inserted {
		list.array[index] = value
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *ArrayList) Add(value int) {
	if list.inserted == len(list.array) {
		list.doubleArray()
	}
	list.array[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= list.inserted {
		if list.inserted == len(list.array) {
			list.doubleArray()
		}

		//Shifting the elements of the array to the right
		for i := list.inserted; i > index; i-- {
			list.array[i] = list.array[i-1]
		}

		list.array[index] = value
		list.inserted++
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}
func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inserted {
		//Shifting the elements of the array to the left
		for i := index; i < list.inserted-1; i++ {
			list.array[i] = list.array[i+1]
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *ArrayList) Reverse() {
	for i := 0; i < list.inserted/2; i++ {
		aux := list.array[i]
		list.array[i] = list.array[(list.inserted-1)-i]
		list.array[(list.inserted-1)-i] = aux
	}
}
