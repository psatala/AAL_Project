package main

import (
	"fmt"
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

func main() {
	rand.Seed(time.Now().UnixNano())

	//generate new problem
	nPartitions := 4
	n := 20000
	maxValue := 10000
	D, Z := generate(n, maxValue)

	//example problem
	//n := 3
	//D := []int{1, 4, 7}
	//Z := []int{10, 13, 5}

	//cyclic rotation for ease of computation
	auxiliary := D[n-1]
	D = D[:n-1]
	D = append([]int{auxiliary}, D...)

	//prefix sum
	for i := 1; i < n; i++ {
		D[i] += D[i-1]
	}

	//comparison between standard and concurrent

	//standard
	start := time.Now()
	r := bruteForce(n, D, Z)
	elapsed := time.Since(start)

	fmt.Printf("Standard:\nIndex: %d\nValue: %d\n", r.index, r.value)
	fmt.Printf("Time:  %f\n", elapsed.Seconds())

	//concurrent
	startConc := time.Now()
	rConc := concurrentBruteForce(n, D, Z, nPartitions)
	elapsedConc := time.Since(startConc)

	fmt.Printf("Concurrent:\nIndex: %d\nValue: %d\n", rConc.index, rConc.value)
	fmt.Printf("Time:  %f\n", elapsedConc.Seconds())

	//linear
	startLin := time.Now()
	rLin := linear(n, D, Z)
	elapsedLin := time.Since(startLin)

	fmt.Printf("Linear:\nIndex: %d\nValue: %d\n", rLin.index, rLin.value)
	fmt.Printf("Time:  %f\n", elapsedLin.Seconds())

}
