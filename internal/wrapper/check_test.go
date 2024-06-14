package wrapper_test

import (
	"testing"

	"github.com/developomp/pompup/internal/wrapper"
)

type RegexMatchCase = struct {
	pattern  string
	s        string
	expected bool
}

func TestRegexMatch(t *testing.T) {
	cases := []RegexMatchCase{
		{"..", "aa", true},
		{".*", "aaa", true},
		{"a\\.b", "a.b", true},
		{"a-b_.*c", "a-b_abc", true},
		{"asdf", "asdf", true},

		// missing letters
		{"asdf", "asd", false},
		{"asdf", "sdf", false},
		{"asd", "asdf", false},
		{"sdf", "asdf", false},

		// case insensitive search
		{"(?i)asdf", "AsDf", true},
	}

	for _, c := range cases {
		result := wrapper.RegexMatch(c.pattern, c.s)

		if result != c.expected {
			t.Fatalf(`RegexMatch("%s", "%s") is %v when it should be %v`, c.pattern, c.s, result, c.expected)
		}
	}
}
