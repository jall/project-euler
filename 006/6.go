package main

import (
    "fmt"
    "github.com/jall/project-euler/numberutil"
)

func main() {
    numbers := numberutil.MakeRange(1, 100)
    difference := absolute(squareOfSums(numbers) - sumOfSquares(numbers))
    fmt.Println(difference)
}

func sumOfSquares(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += (number * number)
    }
    return total
}

func squareOfSums(numbers []int) int {
    sum := numberutil.Sum(numbers)
    return sum * sum
}

func absolute(number int) int {
    if number < 0 {
        return number * -1
    }

    return number
}
