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

func CountingSort(v []int) {
	max := v[0]
	min := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
		if v[i] < min {
			min = v[i]
		}
	}

	len_count := max - min + 1
	count := make([]int, len_count)
	for i := 0; i < len(v); i++ {
		count[v[i]-min]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	ordenado := make([]int, len(v))
	for i := len(v) - 1; i >= 0; i-- {
		indexOrdenado := count[v[i]-min] - 1
		ordenado[indexOrdenado] = v[i]
		count[v[i]-min]--
	}

	copy(v, ordenado)
}

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	fmt.Println(v)

	start_c := time.Now()
	CountingSort(v)
	duration_c := time.Since(start_c)
	fmt.Println("Tempo de execução CountingSort:      ", duration_c)

	fmt.Println(v)
}
