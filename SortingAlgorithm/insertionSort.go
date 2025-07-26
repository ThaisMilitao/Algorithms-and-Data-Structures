package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func creatV(n int, sorted bool) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = int(rand.Intn(100))
	}
	if sorted {
		sort.Ints(v)
	}
	return v
}

func InsertionSort(v []int) {
	for insercao := 1; insercao < len(v); insercao++ {
		temp := v[insercao]
		i := insercao
		for ; i >= 1 && v[i-1] > temp; i-- {
			v[i] = v[i-1]
		}
		v[i] = temp
	}
}

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	fmt.Println(v)

	start_i := time.Now()
	InsertionSort(v)
	duration_i := time.Since(start_i)
	fmt.Println("Tempo de execução Insertion:         ", duration_i)
	fmt.Println(v)
}
