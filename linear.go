package main

import "math"

//calculate distance between two given cities
func calcDist(D []int, roadLength int, a int, b int) int {
	dist := D[b] - D[a]
	if dist < 0 {
		dist += roadLength
	}

	return dist
}

func linear(n int, D []int, Z []int) (r result) {

	r.index = -1
	r.value = math.MaxInt64

	//weird case
	if 0 == n {
		return r
	}

	//variable setup
	zL := 0
	zP := 0
	var kL int64 = 0
	var kP int64 = 0
	roadLength := D[n-1]
	j := 1

	//append copy to make things easier
	D = append(D, D...)
	Z = append(Z, Z...)

	//calculate variables for city 0
	for i := 1; i < n; i++ {
		if calcDist(D, roadLength, 0, i) <= roadLength/2 {
			zL += Z[i]
			kL += int64(Z[i]) * int64(calcDist(D, roadLength, 0, i))
			j++
		} else {
			zP += Z[i]
			kP += int64(Z[i]) * int64(calcDist(D, roadLength, i, 0))
		}
	}

	//store result
	r.index = 0
	r.value = kL + kP

	//main loop
	for i := 0; i < n-1; i++ {
		kL -= int64(calcDist(D, roadLength, i, i+1)) * int64(zL)
		zL -= Z[i+1]
		for calcDist(D, roadLength, i+1, j) <= roadLength/2 {
			zL += Z[j]
			kL += int64(calcDist(D, roadLength, i+1, j)) * int64(Z[j])
			zP -= Z[j]
			kP -= int64(calcDist(D, roadLength, j, i)) * int64(Z[j])
			j++
		}
		zP += Z[i]
		kP += int64(calcDist(D, roadLength, i, i+1)) * int64(zP)

		if kL+kP < r.value {
			r.index = i + 1
			r.value = kL + kP
		}
	}

	return r
}
