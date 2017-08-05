package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
)

func main() {
	maxA := 0
	maxB := 0
	maxN := 0

	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			n := numberOfConsecutivePrimes(a, b)
			if n > maxN {
				maxN = n
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Println("Product: ", maxA*maxB)
	fmt.Println("n: ", maxN)
	fmt.Println("a: ", maxA)
	fmt.Println("b: ", maxB)
}

func numberOfConsecutivePrimes(a, b int) int {
	length := 0

	for n := 0; ; n++ {
		result := (n * n) + (a * n) + b

		if numberutil.IsPrime(result) {
			length++
		} else {
			break
		}
	}

	return length
}
