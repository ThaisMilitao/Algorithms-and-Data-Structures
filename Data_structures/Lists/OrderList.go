// Suponha que você queira criar uma nova implementação do TAD List que sempre se mantém ordenada: OrderedList.
// Uma forma de fazer isso seria anulando a função que permite adicionar em uma posição arbitrária, AddOnIndex,
// e ajustar a implementação de Add(value int) para que ela sempre adicionasse value na posição correta da lista.
// Proveja a implementação das funções de OrderedList, apresentadas na tabela a seguir.
package main

import (
	"errors"
	"fmt"
)

type OrderedList struct {
	array    []int
	inserted int
}

func (list *OrderedList) Init(size int) error {
	if size > 0 {
		list.array = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayList with size <= 0")
	}
}
func (list *OrderedList) DoubleArray() {
	newArray := make([]int, 2*list.inserted)
	for i := 0; i < list.inserted; i++ {
		newArray[i] = list.array[i]
	}
	list.array = newArray
}

func (list *OrderedList) Size() int {
	return list.inserted
}

func (list *OrderedList) bin_search(val int, start int, end int) int {
	if start == end {
		if list.array[start] < val {
			return start + 1
		}
		return start
	}
	mid := (end + start) / 2

	if list.array[mid] == val {
		return mid

	} else if list.array[mid] > val {
		end = (mid - 1)

	} else {
		start = (mid + 1)
	}
	return list.bin_search(val, start, end)
}

func (list *OrderedList) Add(val int) {
	if list.inserted == len(list.array) {
		list.DoubleArray()
	}
	if list.Size() > 0 {
		index := list.bin_search(val, 0, list.inserted-1)
		for i := list.inserted; i > index; i-- {
			list.array[i] = list.array[i-1]
		}
		list.array[index] = val
	} else {
		list.array[0] = val
	}

	list.inserted++
}

func (list *OrderedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < list.inserted {
		for i := index; i < list.inserted-1; i++ {
			list.array[i] = list.array[i+1]
		}
		list.inserted--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func (list *OrderedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.array[index], nil
	} else {
		return -1, errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}
func (list *OrderedList) Set(value, index int) error {
	if index >= 0 && index < list.inserted {
		if value <= list.array[index+1] && value >= list.array[index-1] {
			list.array[index] = value
			return nil
		} else {
			return errors.New(fmt.Sprintf("Invalid input: please insert a value less than %d or greater than %d.", list.array[index+1], list.array[index-1]))
		}
	} else {
		return errors.New(fmt.Sprintf("Invalid index: %d", index))
	}
}

func main() {

}
