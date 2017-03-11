package numberutil

import "math/big"

// Possible combinations of r objects from a set of n objects
// https://en.wikipedia.org/wiki/Binomial_coefficient
func BinomialCoefficient(n, k int) int {
	return FallingFactorial(n, k) / Factorial(k)
}

func BinomialCoefficientBigInt(n, k *big.Int) *big.Int {
	return Div(FallingFactorialBigInt(n, k), FactorialBigInt(k))
}
