package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	amicableNumbers := []int{}
	d := cachedD()

	for a := 1; a < 10000; a++ {
		b := d(a)

		if a != b && a == d(b) {
			amicableNumbers = append(amicableNumbers, a)
		}
	}

	fmt.Println("Sum: ", numberutil.Sum(amicableNumbers))
	fmt.Println("Amicable numbers: ", amicableNumbers)
}

func cachedD() func(n int) int {
	sumOfDivisors := map[int]int{}

	return func(n int) int {
		if _, ok := sumOfDivisors[n]; !ok {
			sumOfDivisors[n] = sumOfProperDivisors(n)
		}

		return sumOfDivisors[n]
	}
}

func sumOfProperDivisors(n int) int {
	primeFactors := numberutil.PrimeFactors(n)
	divisors := []int{}
	for divisor := range numberutil.ProperDivisorsFromPrimeFactors(primeFactors) {
		divisors = append(divisors, divisor)
	}
	return numberutil.Sum(divisors)
}
