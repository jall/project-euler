package main

import "fmt"

func main() {
	var multiples []int

	for x := 0; x < 1000; x++ {
		if (x%3 == 0) || (x%5 == 0) {
			multiples = append(multiples, x)
		}
	}

	fmt.Println("Sum: ", sum(multiples))
}

func sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}
