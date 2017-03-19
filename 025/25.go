package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(1)
	y := big.NewInt(1)
	index := 2

	for {
		// Calculate next value in sequence
		tmp := big.NewInt(0).Add(x, y)
		x = y
		y = tmp

		index++

		// End when number is 1000 digits long
		if len(y.Text(10)) > 1e3-1 {
			break
		}
	}

	fmt.Println("Number:", y)
	fmt.Println("Index:", index)
}
