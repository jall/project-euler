package numberutil

import (
	"reflect"
	"testing"
)

func TestMakeRange(t *testing.T) {
	cases := []struct {
		min  int
		max  int
		want []int
	}{
		{1, 1, []int{1}},
		{0, 5, []int{0, 1, 2, 3, 4, 5}},
		{16, 17, []int{16, 17}},
	}

	for _, c := range cases {
		got := MakeRange(c.min, c.max)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MakeRange(%d, %d) == %d, want %d", c.min, c.max, got, c.want)
		}
	}
}
