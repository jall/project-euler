package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
)

func main() {
	primes := numberutil.SieveOfEratosthenes(2e6)
	fmt.Println(sum(primes))
}

func sum(numbers []int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}
