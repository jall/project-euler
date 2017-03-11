package numberutil

import (
	"reflect"
	"testing"
)

func TestCollatzSequence(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{13, []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
	}

	for _, c := range cases {
		got := CollatzSequence(c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CollatzSequence(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
