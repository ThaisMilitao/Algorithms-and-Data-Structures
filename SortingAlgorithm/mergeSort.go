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

func merge(v []int, esq []int, dir []int) {
	indexEsq, indexDir, indexV := 0, 0, 0
	for indexEsq < len(esq) && indexDir < len(dir) {
		if esq[indexEsq] <= dir[indexDir] {
			v[indexV] = esq[indexEsq]
			indexEsq++
		} else {
			v[indexV] = dir[indexDir]
			indexDir++
		}
		indexV++
	}

	for indexEsq < len(esq) {
		v[indexV] = esq[indexEsq]
		indexEsq++
		indexV++
	}
	for indexDir < len(dir) {
		v[indexV] = dir[indexDir]
		indexDir++
		indexV++
	}
}

func mergeSort(v []int) {
	if len(v) > 1 {
		mid := len(v) / 2

		esq := make([]int, mid)
		copy(esq, v[:mid])

		dir := make([]int, len(v)-len(esq))
		copy(dir, v[mid:])

		mergeSort(esq)
		mergeSort(dir)
		merge(v, esq, dir)
	}
}

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	fmt.Println(v)

	start := time.Now()
	mergeSort(v)
	duration := time.Since(start)
	fmt.Println("Tempo de execução mergeSort:         ", duration)

	fmt.Println(v)

}
