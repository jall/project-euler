package stringutil

import "testing"

func TestLeftPad(t *testing.T) {
	cases := []struct {
		str    string
		pad    string
		length int
		want   string
	}{
		{"1", "0", 4, "0001"},
		{"世界", " ", 10, "        世界"},
		{"", "", 0, ""},
	}
	for _, c := range cases {
		got := LeftPad(c.str, c.pad, c.length)
		if got != c.want {
			t.Errorf("LeftPad(%q, %q, %q) == %q, want %q", c.str, c.pad, c.length, got, c.want)
		}
	}
}
