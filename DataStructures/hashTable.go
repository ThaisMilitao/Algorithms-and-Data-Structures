package main

import (
	"errors"
	"fmt"
)

type Map interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Remove(key int) error
	Size() int
	LoadFactor() float32
	Init()
}

type Tuple struct {
	key   int
	value string
}

var emptyTuple Tuple = Tuple{}

type HashTable struct {
	buckets          [][]Tuple
	elementsInserted int
}

func (table *HashTable) Init(size int) {
	table.buckets = make([][]Tuple, size)
	table.elementsInserted = 0
}

func put(buckets [][]Tuple, key int, value string) {
	bucket := key % len(buckets)
	buckets[bucket] = append(buckets[bucket], Tuple{key, value})
}

func (table *HashTable) increaseBuckets() {
	increasedBuckets := make([][]Tuple, 8*len(table.buckets))
	for i := 0; i < len(table.buckets); i++ {
		for _, tuple := range table.buckets[i] {
			put(increasedBuckets, tuple.key, tuple.value)
		}
	}
	table.buckets = increasedBuckets
}

func (table *HashTable) Put(key int, value string) {
	if table.LoadFactor() > 0.75 {
		table.increaseBuckets()
	}
	put(table.buckets, key, value)
	table.elementsInserted++
}

func (table *HashTable) Get(key int) (string, error) {
	bucket := key % len(table.buckets)
	for _, tuple := range table.buckets[bucket] {
		if tuple.key == key {
			return tuple.value, nil
		}
	}
	return "", errors.New("The table does not contain the specified key!")
}

func (table *HashTable) Remove(key int) error {
	bucket := key % len(table.buckets)
	index := -1
	for i, tuple := range table.buckets[bucket] {
		if tuple.key == key {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("The table does not contain the specified key!")
	} else {
		table.buckets[bucket] = append(table.buckets[bucket][:index], table.buckets[bucket][index+1:]...)
		table.elementsInserted--
		return nil
	}
}

func (table *HashTable) LoadFactor() float32 {
	return float32(table.elementsInserted) / float32(len(table.buckets))
}

func (table *HashTable) Size() int {
	return table.elementsInserted
}

func main() {

	var table HashTable = HashTable{}
	table.Init(4)

	table.Put(1, "um")
	table.Put(2, "dois")
	table.Put(3, "três")
	table.Put(4, "quatro")

	for i := 1; i <= 4; i++ {
		val, err := table.Get(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Key: %d, Value: %s\n", i, val)
		}
	}

	err := table.Remove(2)
	if err != nil {
		fmt.Println("Erro ao remover:", err)
	} else {
		fmt.Println("Elemento com chave 2 removido.")
	}

	_, err = table.Get(2)
	if err != nil {
		fmt.Println("Busca após remoção:", err)
	}

	fmt.Println("Tamanho atual da tabela:", table.Size())

	fmt.Printf("Load factor: %.2f\n", table.LoadFactor())

}
