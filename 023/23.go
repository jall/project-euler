package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

const analyticLimit int = 28123

func main() {
	abundantNumbers := getAbundantNumbers()
	sum := 0

	for i := analyticLimit; i > 0; i-- {
		if isUnableTobeExpressedAsSumOfTwoAbundantNumbers(i, abundantNumbers) {
			sum += i
		}
	}

	fmt.Println("Sum:", sum)
}

func getAbundantNumbers() map[int]int {
	abundantNumbers := make(map[int]int)

	for i := 12; i <= analyticLimit; i++ {
		if numberutil.IsAbundant(i) {
			abundantNumbers[i] = i
		}
	}

	return abundantNumbers
}

func isUnableTobeExpressedAsSumOfTwoAbundantNumbers(i int, abundantNumbers map[int]int) bool {
	for abundant := range abundantNumbers {
		if _, ok := abundantNumbers[i-abundant]; ok {
			return false
		}
	}

	return true
}
