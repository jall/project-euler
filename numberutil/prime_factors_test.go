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

func TestPrimeFactorsGrouped(t *testing.T) {
	cases := []struct {
		in   int
		want map[int]int
	}{
		{20, map[int]int{2:2, 5:1}},
		{600851475143, map[int]int{71:1, 839:1, 1471:1, 6857:1}},
		{1, map[int]int{}},
	}

	for _, c := range cases {
		got := PrimeFactorsGrouped(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PrimeFactorsGrouped(%d) == %v, want %v", c.in, got, c.want)
		}
	}
}
