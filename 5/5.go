package main

import (
    "fmt"
    "github.com/jall/project-euler/numberutil"
)

func main() {
    sharedFactors := make(map[int]int)

    // Compute prime factors for all divisors
    for divisor := range makeRange(1, 20) {
        for factor, count := range numberutil.PrimeFactorsGrouped(divisor) {
            // Figure out shared factors between all divisors
            if (sharedFactors[factor] < count) {
                sharedFactors[factor] = count
            }
        }
    }

    // Multiply shared factors to get smallest number evenly divisible
    // by all divisors
    number := 1
    for factor, exponent := range sharedFactors {
        number = number * numberutil.Pow(factor, exponent)
    }

    fmt.Println(number)
}

func makeRange(min, max int) []int {
    slice := make([]int, max-min+1)
    for i := range slice {
        slice[i] = min + i
    }
    return slice
}
