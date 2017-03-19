package numberutil

import (
	"reflect"
	"testing"
)

func TestExpandFractionToDecimal(t *testing.T) {
	cases := []struct {
		numerator       int
		denominator     int
		wantExpansion   []int
		wantIsRepeating bool
	}{
		{1, 3, []int{3}, true},
		{1, 7, []int{1, 4, 2, 8, 5, 7}, true},
		{1, 8, []int{1, 2, 5}, false},
	}

	for _, c := range cases {
		gotExpansion, gotIsRepeating := ExpandFractionToDecimal(c.numerator, c.denominator)
		if gotIsRepeating != c.wantIsRepeating || !reflect.DeepEqual(gotExpansion, c.wantExpansion) {
			t.Errorf("ExpandFractionToDecimal(%v, %v) == (%v, %v), want (%v, %v)",
				c.numerator,
				c.denominator,
				c.wantExpansion,
				c.wantIsRepeating,
				gotExpansion,
				gotIsRepeating,
			)
		}
	}
}
