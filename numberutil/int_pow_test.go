package numberutil

import "testing"

func TestPow(t *testing.T) {
	cases := []struct {
		num      int
		exponent int
		want     int
	}{
		{1, 1, 1},
		{2, 4, 16},
		{1883, 0, 1},
		{5, 3, 125},
	}

	for _, c := range cases {
		got := Pow(c.num, c.exponent)
		if got != c.want {
			t.Errorf("Pow(%d, %d) == %d, want %d", c.num, c.exponent, got, c.want)
		}
	}
}
