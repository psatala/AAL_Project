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

func generate(n int, maxValue int) ([]int, []int) {
	D := make([]int, n)
	Z := make([]int, n)
	for i := 0; i < n; i++ {
		D[i] = rand.Intn(maxValue)
		Z[i] = rand.Intn(maxValue)
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

//Returns time of bruteForce, concurrentBruteForce, linear, and bool that is true if all solutions are equal and false otherwise
// func solveForAnalysis(D []int, Z []int) (float64, float64, float64, bool)

func main() {
	args := os.Args[1:]
	D := make([]int, 0)
	Z := make([]int, 0)
	var n int = 1000
	var maxValue int = 10000
	var ok bool = true
	var err error
	var step int = 500
	var seed int64 = time.Now().UnixNano()
	var k int = 1
	var r int = 1
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
		modeId = 2
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
	} else {
		for k > 0 {
			for r > 0 {
				D, Z = generate(n, maxValue)
				solveForUser(D, Z)
				r -= 1
			}
			k -= 1
			n += step
		}
	}
}
