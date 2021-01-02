// Epidemia
// Piotr Satala, Piotr Libera

package main

import "math"

func concurrentCalc(n int, D []int, Z []int, begin int, end int, ch chan result) {

	var r result

	r.index = -1
	r.value = math.MaxInt64

	end = int(math.Min(float64(end), float64(n)))

	//main loop
	for i := begin; i < end; i++ {
		temp := calcOneCity(n, D, Z, i, D[n-1])

		if temp < r.value { //found better
			r.index = i
			r.value = temp
		}
	}

	ch <- r
}

func concurrentBruteForce(n int, D []int, Z []int, nPartitions int) (r result) {
	if nPartitions > n/2 {
		nPartitions = n/2
	}
	r.index = -1
	r.value = math.MaxInt64

	partitionSize := n / nPartitions
	ch := make(chan result, nPartitions)

	for i := 0; i < nPartitions; i++ {
		go concurrentCalc(n, D, Z, i*partitionSize, (i+1)*partitionSize, ch)
	}
	for i := 0; i < nPartitions; i++ {
		temp := <-ch
		if temp.value < r.value {
			r.index = temp.index
			r.value = temp.value
		}
	}
	return r
}
