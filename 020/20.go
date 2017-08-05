package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	bigNumber := numberutil.FactorialBigInt(big.NewInt(100))
	sum := 0

	for _, character := range strings.Split(bigNumber.String(), "") {
		digit, _ := strconv.Atoi(character)
		sum += digit
	}

	fmt.Println("Sum of 100!: ", sum)
	fmt.Println("100!: ", bigNumber)
}
