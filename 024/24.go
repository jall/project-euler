package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
	"math"
)

func main() {
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	code := LehmerCode(1e6-1, len(original))
	permutation := LehmerCodeToPermutation(original, code)
	fmt.Println("Lehmer code:", numberutil.IntegersToString(code))
	fmt.Println("Permutation:", numberutil.IntegersToString(permutation))
}

// https://en.wikipedia.org/wiki/Factorial_number_system#Permutations
// https://en.wikipedia.org/wiki/Lehmer_code
func LehmerCode(num, length int) []int {
	if length <= 1 {
		return []int{0}
	}

	multiplier := numberutil.Factorial(length - 1)
	digit := int(math.Floor(float64(num) / float64(multiplier)))

	return append([]int{digit}, LehmerCode(num%multiplier, length-1)...)
}

func LehmerCodeToPermutation(original, code []int) (permutation []int) {
	// Copy the slice so removing elements doesn't break
	set := original[:]
	for _, position := range code {
		permutation = append(permutation, set[position])
		set = splice(set, []int{}, position, 1)
	}
	return permutation
}

func splice(full, part []int, position, length int) []int {
	startHalf := append(full[:position], part...)
	return append(startHalf, full[position+length:]...)
}
