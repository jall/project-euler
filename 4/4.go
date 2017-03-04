package main

import (
	"fmt"
	"strconv"
)

func main() {
	var palindromes []int

	// TODO Work through largest i & j & stop on first match, to avoid scanning
	// all values
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			number := i * j
			if isPalindrome(number) {
				palindromes = append(palindromes, number)
			}
		}
	}

	fmt.Println(max(palindromes))
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	return str == reverse(str)
}

// Reverse returns its argument string reversed rune-wise left to right.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
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
