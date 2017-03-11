package main

import "fmt"

func main() {
	largestN := 1
	largestChain := []int{1}
	largestChainLength := 1

	for n := 1; n < 1e6; n++ {
		sequence := collatzSequence(n)
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

func collatzSequence(n int) []int {
	return collatzSequenceRecursive(n, []int{})
}

func collatzSequenceRecursive(n int, sequence []int) []int {
	sequence = append(sequence, n)

	if n == 1 {
		return sequence
	}

	if n%2 == 0 {
		return collatzSequenceRecursive(n/2, sequence)
	} else {
		return collatzSequenceRecursive(3*n+1, sequence)
	}
}
