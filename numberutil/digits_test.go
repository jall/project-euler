package numberutil

import (
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	cases := []struct {
		in   int
		want []int
	}{
		{0, []int{0}},
		{1, []int{1}},
		{13, []int{1, 3}},
		{7334, []int{7, 3, 3, 4}},
		{912738123, []int{9, 1, 2, 7, 3, 8, 1, 2, 3}},
	}

	for _, c := range cases {
		got := Digits(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Digits(%d) == %v, want %v", c.in, got, c.want)
		}
	}
}
