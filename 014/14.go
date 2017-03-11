package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	largestN := 1
	largestChain := []int{1}
	largestChainLength := 1

	for n := 1; n < 1e6; n++ {
		sequence := numberutil.CollatzSequence(n)
		length := len(sequence)

		if length > largestChainLength {
			largestN = n
			largestChain = sequence
			largestChainLength = length
		}
	}

	fmt.Println("n: ", largestN)
	fmt.Println("Chain length: ", largestChainLength)
	fmt.Println("Chain: ", largestChain)
}
