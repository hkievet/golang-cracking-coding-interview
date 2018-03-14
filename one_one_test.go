package interview

import "testing"

func TestHasUniqueCharacters(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abc", true},
		{"abcdefga", false},
		{"hello goodbye", false},
		{"foo bar", false},
		{"fu bar", true},
	}

	for _, c := range cases {
		result := HasUniqueCharacters(c.in)
		if result != c.want {
			t.Errorf("HasUniqueCharacters(%s), want %t, got %t", c.in, c.want, result)
		}
	}
}
