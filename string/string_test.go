package string

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		value1, value2, want string
	}{
		{"Hello, new kevin!", "I'm cute.", "Hello, new kevin!I'm cute."},
		{"", "", ""},
	}

	for _, c := range tests {
		got := Concat(c.value1, c.value2)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.value1, c.value2, got, c.want)
		}
	}	
}