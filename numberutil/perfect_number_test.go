package numberutil

import "testing"

func TestAliquotSum(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{6, 6},
		{28, 28},
		{496, 496},
		{1, 0},
		{12, 16},
		{15, 9},
	}

	for _, c := range cases {
		got := AliquotSum(c.in)
		if got != c.want {
			t.Errorf("AliquotSum(%d) == %d, want %q", c.in, got, c.want)
		}
	}
}

func TestIsPerfect(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{6, true},
		{28, true},
		{496, true},
		{1, false},
		{495, false},
		{497, false},
		{253718236, false},
	}

	for _, c := range cases {
		got := IsPerfect(c.in)
		if got != c.want {
			t.Errorf("IsPerfect(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsAbundant(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{6, false},
		{28, false},
		{496, false},
		{12, true},
		{18, true},
		{120, true},
		{15, false},
		{43, false},
	}

	for _, c := range cases {
		got := IsAbundant(c.in)
		if got != c.want {
			t.Errorf("IsAbundant(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsDeficient(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, true},
		{2, true},
		{21, true},
		{25, true},
		{6, false},
		{28, false},
		{496, false},
		{12, false},
		{120, false},
	}

	for _, c := range cases {
		got := IsDeficient(c.in)
		if got != c.want {
			t.Errorf("IsDeficient(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}
