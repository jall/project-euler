package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	for n := 1; n < 1e8; n++ {
		triangleNumber := numberutil.SumOfNConsecutiveNumbers(n)
		primeFactors := numberutil.PrimeFactors(triangleNumber)
		factors := factorsFromPrimeFactors(primeFactors)
		if len(factors) > 500 {
			fmt.Println("n: ", n)
			fmt.Println("Sum: ", triangleNumber)
			break
		}
	}
}

func factorsFromPrimeFactors(primeFactors []int) map[int]int {
	factors := map[int]int{}

	for _, set := range numberutil.PowerSetWithOptions(primeFactors, true, false) {
		factor := 1
		for _, number := range set {
			factor *= number
		}
		factors[factor]++
	}

	return factors
}
