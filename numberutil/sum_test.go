package numberutil

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		numbers     []int
		want     int
	}{
		{[]int{1}, 1},
		{[]int{2, 4, 16}, 22},
		{[]int{0, 4, 4, 4, 0, 2, 1}, 15},
	}

	for _, c := range cases {
		got := Sum(c.numbers)
		if got != c.want {
			t.Errorf("Sum(%v) == %d, want %d", c.numbers, got, c.want)
		}
	}
}
