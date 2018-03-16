package interview

import "testing"

func TestEscapeSpaces(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"hello world", "hello%20world"},
		{"  blam  ", "%20%20blam%20%20"},
	}

	for _, c := range cases {
		result := EscapeSpaces(c.in)
		if result != c.want {
			t.Errorf("EscapeSpaces(%s), want %s, got %s", c.in, c.want, result)
		}
	}
}
