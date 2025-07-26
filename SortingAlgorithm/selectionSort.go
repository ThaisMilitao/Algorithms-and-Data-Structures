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

func SelectionSort(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++{
		iMenor := varredura
		for i:=varredura+1; i < len(v); i++{
			if v[i] < v[iMenor] {
				iMenor = i
			}
		}
	v[varredura],v[iMenor] = v[iMenor],v[varredura]  
	}
}

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	fmt.Println(v)	
	
	start_s := time.Now()
	SelectionSort(v)
	duration_s := time.Since(start_s)
	fmt.Println("Tempo de execução SelectionSort:     ",duration_s)
	
	fmt.Println(v)
}