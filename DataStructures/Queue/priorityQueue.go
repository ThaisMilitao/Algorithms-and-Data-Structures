package main

import "fmt"

// heap de minimo
type PQ interface {
	Add(value int)
	Remove(value int)
	Pool() int
}

type BinaryHeap struct {
	arr []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{arr: make([]int, 0)}
}

func (h *BinaryHeap) Add(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *BinaryHeap) Remove(value int) {
	index := -1
	for i, v := range h.arr {
		if v == value {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	lastIndex := len(h.arr) - 1
	h.arr[index] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]

	if index < len(h.arr) {
		h.heapifyUp(index)
		h.heapifyDown(index)
	}
}

func (h *BinaryHeap) Pool() int {
	if len(h.arr) == 0 {
		panic("heap is empty")
	}

	min := h.arr[0]
	h.Remove(min)
	return min
}

func (h *BinaryHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.arr[index] >= h.arr[parentIndex] {
			break
		}
		h.arr[index], h.arr[parentIndex] = h.arr[parentIndex], h.arr[index]
		index = parentIndex
	}
}

func (h *BinaryHeap) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild < len(h.arr) && h.arr[leftChild] < h.arr[smallest] {
		smallest = leftChild
	}

	if rightChild < len(h.arr) && h.arr[rightChild] < h.arr[smallest] {
		smallest = rightChild
	}

	if smallest != index {
		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		h.heapifyDown(smallest)
	}
}

func main() {
	heap := NewBinaryHeap()

	heap.Add(5)
	heap.Add(3)
	heap.Add(8)
	heap.Add(1)
	heap.Add(2)
	
	fmt.Println("Pool:", heap.Pool())
	fmt.Println("Pool:", heap.Pool())
	fmt.Println("Pool:", heap.Pool())

	heap.Add(4)
	heap.Remove(5)

	fmt.Println("Pool:", heap.Pool())
	fmt.Println("Pool:", heap.Pool())
}
