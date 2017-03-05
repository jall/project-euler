package main

import (
	"fmt"
)

// This search is very greedy & could find the answer much quicker, but I
// had fun implementing the PythagoreanTriple generator & coprime checker.
func main() {
	// Tweak this function's parameters to generate a larger set of triples
	triples := generatePythagoreanTriples(10, 50)
	isFound := false

	for _, triple := range triples {
		if (triple.a + triple.b + triple.c) == 1000 {
			fmt.Println("Success: ", triple)
			fmt.Println("Product: ", (triple.a * triple.b * triple.c))
			isFound = true
			break
		}
	}

	if !isFound {
		fmt.Println("Failure; checked these triples: ", triples)
	}
}

type PythagoreanTriple struct {
	a int
	b int
	c int
}

// Uses Euclid's formula
// https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
// TODO Figure out more informative limit parameters
func generatePythagoreanTriples(limit, kLimit int) []PythagoreanTriple {
	var triples []PythagoreanTriple
	var triple PythagoreanTriple
	primitives := generatePrimitivePythagoreanTriples(limit)

	for k := 1; k <= kLimit; k++ {
		for _, primitive := range primitives {
			triple = PythagoreanTriple{
				primitive.a * k,
				primitive.b * k,
				primitive.c * k,
			}
			triples = append(triples, triple)
		}
	}

	return triples
}

// Uses Euclid's formula
// https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
// This version only output primitive triples (a, b, c are coprime)
// TODO Figure out more informative limit parameter
func generatePrimitivePythagoreanTriples(limit int) []PythagoreanTriple {
	var triples []PythagoreanTriple

	// n & m must be greater than 0
	for n := 1; n < limit; n++ {
		// m must be greater than n
		for m := n + 1; m < limit; m++ {
			// Restrict to primitive triples
			if isCoprime(m, n) && !(isOdd(m) && isOdd(n)) {
				a := (m * m) - (n * n)
				b := 2 * m * n
				c := (m * m) + (n * n)
				triples = append(triples, PythagoreanTriple{a, b, c})
			}
		}
	}

	return triples
}

func isOdd(x int) bool {
	return x%2 != 0
}

func isCoprime(a, b int) bool {
	return greatestCommonDivisor(a, b) == 1
}

// Extended Euclidean algorithm
// https://en.wikibooks.org/wiki/Algorithm_Implementation/Mathematics/Extended_Euclidean_algorithm
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
