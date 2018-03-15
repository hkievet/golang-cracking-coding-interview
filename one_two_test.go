package interview

import "testing"

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"abc", "bca", true},
		{"hello", "ello", false},
		{"hello", "wasuh", false},
		{"a e i o u", "l m n p q", false},
		{"a e i o u", "i o e a u p", false},
		{"a e i o u", "i o e a u", true},
	}

	for _, c := range cases {
		result := IsPermutation(c.in1, c.in2)
		if result != c.want {
			t.Errorf("IsPermutation(%s, %s), want %t, got %t", c.in1, c.in2, c.want, result)
		}
	}
}
