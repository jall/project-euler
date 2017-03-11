package numberutil

import "testing"

func TestSumOfNConsecutiveNumbers(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 3},
		{57, 1653},
		{100, 5050},
	}

	for _, c := range cases {
		got := SumOfNConsecutiveNumbers(c.n)
		if got != c.want {
			t.Errorf("SumOfNConsecutiveNumbers(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}
