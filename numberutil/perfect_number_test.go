package numberutil

import "testing"

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
			t.Errorf("IsPerfect(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
