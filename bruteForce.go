// Epidemia
// Piotr Satala, Piotr Libera

package main

import "math"

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

func bruteForce(n int, D []int, Z []int) (r result) {

	r.index = -1
	r.value = math.MaxInt64

	//weird case
	if n < 1 {
		return r
	}

	//main loop
	for i := 0; i < n; i++ {
		temp := calcOneCity(n, D, Z, i, D[n-1])

		if temp < r.value { //found better
			r.index = i
			r.value = temp
		}
	}

	return r
}
