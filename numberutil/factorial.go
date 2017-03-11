package numberutil

import "math/big"

// TODO Tests for non-bigInts

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

// https://en.wikipedia.org/wiki/Falling_and_rising_factorials
func FallingFactorial(x, n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= (x - i + 1)
	}
	return factorial
}

func FactorialBigInt(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(one) == 0 {
		return one
	} else {
		return Mul(n, FactorialBigInt(Sub(n, one)))
	}
}

// https://en.wikipedia.org/wiki/Falling_and_rising_factorials
func FallingFactorialBigInt(x, n *big.Int) *big.Int {
	factorial := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(n) != 1; i.Add(i, big.NewInt(1)) {
		factorial = Mul(factorial, Add(Sub(x, i), big.NewInt(1)))
	}
	return factorial
}
