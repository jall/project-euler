package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	primes := numberutil.SieveOfEratosthenes(1e6)
	primeCount := len(primes)
	fmt.Println("Number of primes: ", primeCount)
	if primeCount > 10000 {
		fmt.Println("10001st prime: ", primes[10000])
	}
}
