package numberutil

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{20, []int{2, 2, 5}},
		{600851475143, []int{71, 839, 1471, 6857}},
		{1, []int{}},
	}

	for _, c := range cases {
		got := PrimeFactors(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PrimeFactors(%d) == %v, want %v", c.in, got, c.want)
		}
	}
}
