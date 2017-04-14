package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int64
	var b int64
	set := make(map[string]bool)

	for a = 2; a <= 100; a++ {
		for b = 2; b <= 100; b++ {
			power := Pow(big.NewInt(a), big.NewInt(b)).String()
			set[power] = true
		}
	}

	fmt.Println("Number: ", len(set))
}

func Pow(x, y *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, nil)
}
