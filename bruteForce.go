package main

import "math"

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
