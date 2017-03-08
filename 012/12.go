package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
	"strconv"
	"strings"
)

func main() {
	for n := 1; n < 1e8; n++ {
		triangleNumber := sumOfNConsecutiveNumbers(n)
		primeFactors := numberutil.PrimeFactors(triangleNumber)
		factors := factorsFromPrimeFactors(primeFactors)
		if len(factors) > 500 {
			fmt.Println(n)
			fmt.Println(triangleNumber)
			fmt.Println(factors)
			break
		}
	}
}

func factorsFromPrimeFactors(primeFactors []int) map[int]int {
	factors := map[int]int{}

	for _, set := range powerSet(primeFactors) {
		factor := 1
		for _, number := range set {
			factor *= number
		}
		factors[factor]++
	}

	return factors
}

// Lovely algorithm from Arturo Magidin on StackExchange
// http://math.stackexchange.com/a/92087/377473
func powerSet(set []int) [][]int {
	var powerSet [][]int
	cardinality := len(set)

	// Trivial subsets:
	// 	 Starting from 0 includes the empty set
	//   Ending at 2^k includes the original set
	// To avoid both, go from 1 => 2^k -1
	for i := 0; i < numberutil.Pow(2, cardinality); i++ {
		var subSet []int
		binary := leftPad(strconv.FormatInt(int64(i), 2), "0", cardinality)
		for key, rune := range strings.Split(binary, "") {
			bit, _ := strconv.Atoi(rune)
			if bit == 1 {
				subSet = append(subSet, set[key])
			}
		}
		powerSet = append(powerSet, subSet)
	}

	return powerSet
}

func leftPad(str string, pad string, length int) string {
	for {
		if len(str) >= length {
			return str
		}
		str = pad + str
	}
}

// Clever old (/young) Gauss
func sumOfNConsecutiveNumbers(n int) int {
	return (n * (n + 1)) / 2
}
