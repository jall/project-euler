package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	fmt.Println(numberutil.Max(numberutil.PrimeFactors(600851475143)))
}
