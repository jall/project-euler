package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

// https://en.wikipedia.org/wiki/Narcissistic_number
//
// Skip the single digit numbers because they 'aren't a sum'.
//
// Max is because 6 digits is the maximum length before the sum of digit powers stops growing fast enough to
// reach the digit count required, and 6*(9^5) = 354294.
//
// The highest output for any individual digit is 9^5 = 59049, and a brief experiment shows that only numbers
// up to 6 digits are possible, as the largest possible 7 digit number score is 7*(9^5) = 413343, which is
// only 6 digits long.
func main() {
	narcissists := []int{}

	for i := 10; i <= 354294; i++ {
		sum := 0

		for _, digit := range numberutil.Digits(i) {
			sum += numberutil.Pow(digit, 5)
		}

		if sum == i {
			narcissists = append(narcissists, i)
		}
	}

	fmt.Println("Narcissists: ", narcissists)
	fmt.Println("Sum: ", numberutil.Sum(narcissists))
}
