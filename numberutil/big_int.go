package numberutil

import "math/big"

// TODO Tests

func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func Div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}
