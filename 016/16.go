package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	number := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	sumOfDigits := 0

	for _, str := range strings.Split(number.String(), "") {
		digit, _ := strconv.Atoi(str)
		sumOfDigits += digit
	}

	fmt.Println("2^1000: ", number)
	fmt.Println("Sum of digits: ", sumOfDigits)
}
