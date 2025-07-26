package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)
func creatV(n int,sorted bool)[]int{
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = int(rand.Intn(100))
	}
	if sorted {
		sort.Ints(v)
	}
	return v
}

func partitionPivotFim(v []int, ini int, fim int) int{
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++{
		if v[i] <= pivot{
			v[pIndex],v[i] = v[i],v[pIndex]
			pIndex++
		}
	}
	v[pIndex],v[fim] = v[fim],v[pIndex]
	return pIndex
}

func quickSortPivotFim(v []int,ini int,fim int){
	if ini < fim{
		pivot := partitionPivotFim(v,ini,fim)
		quickSortPivotFim(v,ini,pivot-1)
		quickSortPivotFim(v,pivot+1,fim)
	}
}

// quicksort com pivot aleatorio
func partition(v []int, ini int, fim int) int{
	index_pivot := rand.Intn(fim-ini+1) + ini
	v[index_pivot], v[fim] = v[fim], v[index_pivot]
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++{
		if v[i] <= pivot{
			v[pIndex],v[i] = v[i],v[pIndex]
			pIndex++
		}
	}
	v[pIndex],v[fim] = v[fim],v[pIndex]
	return pIndex
}

func quickSort(v []int,ini int,fim int){
	if ini < fim{
		pivot := partition(v,ini,fim)
		quickSort(v,ini,pivot-1)
		quickSort(v,pivot+1,fim)
	}
}

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	v_quicksortSemRand := make([]int, len(v))
	copy(v_quicksortSemRand, v)

	v_quicksort := make([]int, len(v))
	copy(v_quicksort, v)	

	start_q := time.Now()
	quickSortPivotFim(v_quicksortSemRand,0,len(v_quicksortSemRand)-1)
	duration_q := time.Since(start_q)
	fmt.Println("Tempo de execução quickSort sem rand:",duration_q)
	// fmt.Println(v_quicksortSemRand)

	start_qs := time.Now()
	quickSort(v_quicksort,0,len(v_quicksort)-1)
	duration_qs := time.Since(start_qs)
	fmt.Println("Tempo de execução quickSort com rand:",duration_qs)
	// fmt.Println(v_quicksort)
}