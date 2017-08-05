package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
	"math/big"
)

func main() {
	m, n := big.NewInt(20), big.NewInt(20)

	fmt.Println("Routes:", numberutil.BinomialCoefficientBigInt(numberutil.Add(m, n), n))
}
