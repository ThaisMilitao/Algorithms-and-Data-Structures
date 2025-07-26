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

func main() {
	v := creatV(int(math.Pow(10, 5)), true)
	// v := creatV(int(math.Pow(10, 5)), false)

	fmt.Println(v)

	start_b := time.Now()
	BubbleSort(v)
	duration_b := time.Since(start_b)
	fmt.Println("Tempo de execução BubbleSort:        ", duration_b)

	fmt.Println(v)
}
