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

func SelectionSort(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		iMenor := varredura
		for i := varredura + 1; i < len(v); i++ {
			if v[i] < v[iMenor] {
				iMenor = i
			}
		}
		v[varredura], v[iMenor] = v[iMenor], v[varredura]
	}
}

func BubbleSort(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ {
		trocou := false
		for i := 0; i < len(v)-1-varredura; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				trocou = true
			}
		}
		if !trocou {
			break
		}
	}
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

func partitionPivotFim(v []int, ini int, fim int) int {
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		}
	}
	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}

func quickSortPivotFim(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partitionPivotFim(v, ini, fim)
		quickSortPivotFim(v, ini, pivot-1)
		quickSortPivotFim(v, pivot+1, fim)
	}
}

// quicksort com pivot aleatorio
func partition(v []int, ini int, fim int) int {
	index_pivot := rand.Intn(fim-ini+1) + ini
	v[index_pivot], v[fim] = v[fim], v[index_pivot]
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		}
	}
	v[pIndex], v[fim] = v[fim], v[pIndex]
	return pIndex
}

func quickSort(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partition(v, ini, fim)
		quickSort(v, ini, pivot-1)
		quickSort(v, pivot+1, fim)
	}
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

	v_SelectionSort := make([]int, len(v))
	copy(v_SelectionSort, v)

	v_BubbleSort := make([]int, len(v))
	copy(v_BubbleSort, v)

	v_InsertionSort := make([]int, len(v))
	copy(v_InsertionSort, v)

	v_mergeSort := make([]int, len(v))
	copy(v_mergeSort, v)

	v_quicksortSemRand := make([]int, len(v))
	copy(v_quicksortSemRand, v)

	v_quicksort := make([]int, len(v))
	copy(v_quicksort, v)

	v_CountingSort := make([]int, len(v))
	copy(v_CountingSort, v)

	// fmt.Println(v)

	start_s := time.Now()
	SelectionSort(v_SelectionSort)
	duration_s := time.Since(start_s)
	fmt.Println("Tempo de execução SelectionSort:     ", duration_s)
	// fmt.Println(v_SelectionSort)

	start_b := time.Now()
	BubbleSort(v_BubbleSort)
	duration_b := time.Since(start_b)
	fmt.Println("Tempo de execução BubbleSort:        ", duration_b)
	// fmt.Println(v_BubbleSort)

	start_i := time.Now()
	InsertionSort(v_InsertionSort)
	duration_i := time.Since(start_i)
	fmt.Println("Tempo de execução Insertion:         ", duration_i)
	// fmt.Println(v_InsertionSort)

	start := time.Now()
	mergeSort(v_mergeSort)
	duration := time.Since(start)
	fmt.Println("Tempo de execução mergeSort:         ", duration)
	// fmt.Println(v_mergeSort)

	start_q := time.Now()
	quickSortPivotFim(v_quicksortSemRand, 0, len(v_quicksortSemRand)-1)
	duration_q := time.Since(start_q)
	fmt.Println("Tempo de execução quickSort sem rand:", duration_q)
	// fmt.Println(v_quicksortSemRand)

	start_qs := time.Now()
	quickSort(v_quicksort, 0, len(v_quicksort)-1)
	duration_qs := time.Since(start_qs)
	fmt.Println("Tempo de execução quickSort com rand:", duration_qs)
	// fmt.Println(v_quicksort)

	start_c := time.Now()
	CountingSort(v_CountingSort)
	duration_c := time.Since(start_c)
	fmt.Println("Tempo de execução CountingSort:      ", duration_c)
	// fmt.Println(v_CountingSort)
}
