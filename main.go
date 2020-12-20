package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type result struct {
	index int
	value int64
}

func generate(n int, maxValue int) ([]int, []int) {
	D := make([]int, n)
	Z := make([]int, n)
	for i := 0; i < n; i++ {
		D[i] = rand.Intn(maxValue)
		Z[i] = rand.Intn(maxValue)
	}
	return D, Z
}

func calcOneCity(n int, D []int, Z []int, index int, roadLength int) int64 {
	var total int64 = 0
	for i := 0; i < n; i++ {
		dist := D[index] - D[i]
		if dist < 0 {
			dist = -dist
		}
		dist = int(math.Min(float64(dist), float64(roadLength-dist)))
		total += (int64(dist) * int64(Z[i]))
	}
	return total
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nPartitions := 4
	n := 40000
	maxValue := 10000
	D, Z := generate(n, maxValue)

	//n := 3
	//D := []int{7, 1, 4}
	//Z := []int{10, 13, 5}

	//prefix sum
	for i := 1; i < n; i++ {
		D[i] += D[i-1]
	}

	//comparison between standard and concurrent

	//standard
	start := time.Now()
	r := bruteForce(n, D, Z)
	elapsed := time.Since(start)

	fmt.Printf("%d\n%d\n", r.index, r.value)
	fmt.Printf("%f\n", elapsed.Seconds())

	//concurrent
	startConc := time.Now()
	rConc := concurrentBruteForce(n, D, Z, nPartitions)
	elapsedConc := time.Since(startConc)

	fmt.Printf("%d\n%d\n", rConc.index, rConc.value)
	fmt.Printf("%f\n", elapsedConc.Seconds())

}
