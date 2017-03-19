package numberutil

import "testing"

func TestIntegersToString(t *testing.T) {
	cases := []struct {
		numbers []int
		want    string
	}{
		{[]int{1, 2, 3, 4, 5}, "12345"},
		{[]int{0, 0, 0, 1}, "0001"},
		{[]int{9, 0}, "90"},
	}

	for _, c := range cases {
		got := IntegersToString(c.numbers)
		if got != c.want {
			t.Errorf("IntegersToString(%v) == %v, want %v", c.numbers, got, c.want)
		}
	}
}
