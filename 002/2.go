package main

import (
	"fmt"
)

func main() {
    total := 0

    // Begin Fibonacci sequence with 1 & 2
    x := 1
    y := 2

    for {
        // If latest value is even, add to total
        if (y%2 == 0) {
            total = total + y
        }

        // Calculate next value in sequence
        tmp := x
        x = y
        y = x + tmp

        // Only consider cases under 4 million
        if (y > 4e6) {
            break;
        }
    }

    fmt.Println("Answer: ", total)
}
