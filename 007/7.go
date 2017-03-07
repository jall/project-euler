package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieveOfEratosthenes(1e6)
	primeCount := len(primes)
	fmt.Println("Number of primes: ", primeCount)
	if primeCount > 10000 {
		fmt.Println("10001st prime: ", primes[10000])
	}
}

// This version of the algorithm is from wikipedia
// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func sieveOfEratosthenes(limit int) []int {
	candidates := makeBoolArray(limit)

	// Zero and 1 are not primes
	candidates[0] = false
	candidates[1] = false

	for i := 2; i < int(math.Sqrt(float64(limit))); i++ {
		if candidates[i] {
			for j := (i * i); j < limit; j += i {
				candidates[j] = false
			}
		}
	}

	var primes []int

	for number, isPrime := range candidates {
		if isPrime {
			primes = append(primes, number)
		}
	}

	return primes
}

func makeBoolArray(size int) []bool {
	booleans := make([]bool, size)

	for i := range booleans {
		booleans[i] = true
	}

	return booleans
}
