package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	for n := 1; n < 1e8; n++ {
		triangleNumber := numberutil.SumOfNConsecutiveNumbers(n)
		primeFactors := numberutil.PrimeFactors(triangleNumber)
		divisors := numberutil.ProperDivisorsFromPrimeFactors(primeFactors)
		if len(divisors) > 500 {
			fmt.Println("n: ", n)
			fmt.Println("Sum: ", triangleNumber)
			break
		}
	}
}
