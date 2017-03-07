package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
	"github.com/jall/project-euler/stringutil"
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

	fmt.Println(numberutil.Max(palindromes))
}

func isPalindrome(number int) bool {
	str := strconv.Itoa(number)
	return str == stringutil.Reverse(str)
}
