// Epidemia
// Piotr Satala, Piotr Libera

package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)

type result struct {
	index int
	value int64
}

// Przewidywana zlozonosc T(n)
func T (n int, analyzedAlgorithm int) int {
	if analyzedAlgorithm > 0{
		return n*n
	}
	return n
}

func generate(n int, maxValue int, maxD int) ([]int, []int) {
	D := make([]int, n)
	Z := make([]int, n)
	if maxD >= 0 {
		for i := 0; i < n; i++ {
			D[i] = rand.Intn(maxD)
			if D[i] == 0 {
				D[i] = 1
			}
			Z[i] = rand.Intn(maxValue)
		}
		D[n-1] = rand.Intn(maxValue)
	} else {
		for i := 0; i < n; i++ {
			D[i] = rand.Intn(maxValue)
			if D[i] == 0 {
				D[i] = 1
			}
			Z[i] = rand.Intn(maxValue)
		}
	}
	return D, Z
}

func solveForUser(D []int, Z []int) {
	nPartitions := 4
	n := len(D)

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
	result := bruteForce(n, D, Z)
	elapsed := time.Since(start)

	fmt.Printf("Standard:\nIndex: %d\nValue: %d\n", result.index, result.value)
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

func solveForAnalysis(D []int, Z []int, analyzedAlgorithm int) time.Duration {
	n := len(D)
	//cyclic rotation for ease of computation
	auxiliary := D[n-1]
	D = D[:n-1]
	D = append([]int{auxiliary}, D...)
	//prefix sum
	for i := 1; i < n; i++ {
		D[i] += D[i-1]
	}

	start := time.Now()
	algorithmTime := time.Since(start)
	if analyzedAlgorithm == 1 {
		bruteForce(n, D, Z)
		algorithmTime = time.Since(start)
	} else if analyzedAlgorithm == 2 {
		concurrentBruteForce(n, D, Z, 4)
		algorithmTime = time.Since(start)
	} else {
		resultLinear := linear(n, D, Z)
		algorithmTime = time.Since(start)
		if n <= 1000 {
			resultBrute := bruteForce(n, D, Z)
			if resultLinear.index != resultBrute.index || resultLinear.value != resultBrute.value {
				fmt.Printf("ERROR: Linear and brute solutions not equal\n")
			}
		}
	}
	return algorithmTime
}

func main() {
	args := os.Args[1:]
	D := make([]int, 0)
	Z := make([]int, 0)
	var n int = 1000
	var maxValue int = 10000
	var analyzedAlgorithm int = 0
	var ok bool = true
	var err error
	var step int = 500
	var seed int64 = time.Now().UnixNano()
	var k int = 1
	var r int = 1
	var maxD int = -1
	var modeId int = 1
	if len(args) < 1 {
		printHelp()
		return
	}
	if args[0] == "-m1" {
		modeId = 1
		D, Z, ok = getFullInput()
		n = len(D)
		if !ok {
			fmt.Printf("Blad wczytywania danych\n")
			printHelp()
			return
		}
	} else if args[0] == "-m2" || args[0] == "-m3" {
		if args[0] == "-m2" {
			modeId = 2
		} else {
			modeId = 3
		}
		for i := 1; i < len(args); i += 1 {
			switch args[i][:2] {
			case "-n":
				n, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr n\n")
					return
				}
			case "-w":
				maxValue, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr w\n")
					return
				}
			case "-k":
				k, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr k\n")
					return
				}
			case "-g":
				analyzedAlgorithm, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr g\n")
					return
				}
			case "-c":
				maxD, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr c\n")
					return
				}
			case "-t":
				step, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr t\n")
					return
				}
			case "-s":
				var tempSeed int
				tempSeed, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr s\n")
					return
				}
				seed = int64(tempSeed)
			case "-r":
				r, err = strconv.Atoi(args[i][2:])
				if err != nil {
					fmt.Printf("Bledny parametr r\n")
					return
				}
			}
		}
		rand.Seed(seed)
	} else {
		printHelp()
		return
	}

	if modeId == 1 {
		solveForUser(D, Z)
	} else if modeId == 2{
		D, Z = generate(n, maxValue, maxD)
		solveForUser(D, Z)
	} else {
		linearTimes := make([]int, k)
		var lt time.Duration
		kCount := 0
		for kCount < k {
			rCount := r
			var lts time.Duration
			for rCount > 0 {
				D, Z = generate(n, maxValue, maxD)
				lt= solveForAnalysis(D, Z, analyzedAlgorithm)
				if !ok {
					fmt.Printf("Wyniki metody brutalnej i liniowej byly rozne")
					return
				}
				lts += lt
				rCount -= 1
			}
			linearTimes[kCount] =int(lts.Nanoseconds()) / r
			kCount += 1
			n += step
		}
		n -= step*k
		c := float64(linearTimes[k/2]) / float64(T(n+step*(k/2), analyzedAlgorithm))
		fmt.Printf("n \t t(n)[ms] \t q(n)\n")
		for i := 0; i < k; i += 1 {
			fmt.Printf("%d \t %f \t %f\n", n+step*i, float64(linearTimes[i])/1e6, float64(linearTimes[i])/(c*float64(T(n+step*i, analyzedAlgorithm))))
		}
	}
}
