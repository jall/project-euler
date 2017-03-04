package main

import (
	"fmt"
	"github.com/jall/project-euler/numberutil"
)

func main() {
	var multiples []int

	for x := 0; x < 1000; x++ {
		if (x%3 == 0) || (x%5 == 0) {
			multiples = append(multiples, x)
		}
	}

	fmt.Println("Sum: ", numberutil.Sum(multiples))
}
