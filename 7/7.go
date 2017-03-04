package main

import (
    "fmt"
    "github.com/jall/project-euler/numberutil"
    "math"
    "sort"
)

func main() {
    primes := sieveOfEratosthenes(2e5)
    sort.Ints(primes)
    fmt.Println("Number of primes: ", len(primes))
    fmt.Println("10001st prime: ", primes[10001 - 1])
}

func sieveOfEratosthenes(limit int) []int {
    candidates := numberutil.MakeRange(2, limit)

    for prime:= 2; prime < int(math.Sqrt(float64(limit))); prime++ {
        candidates = filterMultiplesFromSlice(candidates, prime)
    }

    return candidates
}

func filterMultiplesFromSlice(slice []int, factor int) []int {
    filtered := make([]int, 0)

    for _, number := range slice {
        if (number % factor != 0 || number == factor) {
            filtered = append(filtered, number)
        }
    }

    return filtered
}

// Removes element from slice without preserving order
// http://stackoverflow.com/a/37335777/7002606
func remove(slice []int, i int) []int {
    // Swap element to remove with last element.
    slice[len(slice)-1], slice[i] = slice[i], slice[len(slice)-1]
    return slice[:len(slice)-1]
}
