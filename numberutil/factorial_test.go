package numberutil

import (
	"math/big"
	"testing"
)

func TestFactorialBigInt(t *testing.T) {
	cases := []struct {
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(1), getBigInt("1")},
		{big.NewInt(6), getBigInt("720")},
		{big.NewInt(11), getBigInt("39916800")},
		{big.NewInt(14), getBigInt("87178291200")},
	}

	for _, c := range cases {
		got := FactorialBigInt(c.n)
		if c.want.Cmp(got) != 0 {
			t.Errorf("FactorialBigInt(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}

func TestFallingFactorialBigInt(t *testing.T) {
	cases := []struct {
		x    *big.Int
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(10), big.NewInt(4), big.NewInt(5040)},
		{big.NewInt(21), big.NewInt(11), getBigInt("14079294028800")},
	}

	for _, c := range cases {
		got := FallingFactorialBigInt(c.x, c.n)
		if c.want.Cmp(got) != 0 {
			t.Errorf("FallingFactorialBigInt(%d, %d) == %d, want %d", c.x, c.n, got, c.want)
		}
	}
}

func getBigInt(str string) *big.Int {
	tmp := big.NewInt(0)
	num, _ := tmp.SetString(str, 10)
	return num
}
