package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(primeFactors(600851475143)))
}

func primeFactors(number int) []int {
	var factors []int
	divisor := 2

	for number > 1 {
		for number%divisor == 0 {
			factors = append(factors, divisor)
			number = number / divisor
		}
		divisor = divisor + 1
	}

	return factors
}

func max(numbers []int) int {
	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
