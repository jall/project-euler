package numberutil

import "testing"

func TestMax(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 4, 1e5, 5, 199, 123456}, 123456},
		{[]int{5, 4, 3, 2, 1}, 5},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{0}, 0},
	}

	for _, c := range cases {
		got := Max(c.in)
		if got != c.want {
			t.Errorf("Max(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
