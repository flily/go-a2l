package strutil

import (
	"testing"

	"strings"
)

func TestCommentInLine(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"This is a comment.", "/* This is a comment. */"},
	}

	for _, test := range tests {
		result := CommentInLine(test.input)
		if result != test.expected {
			t.Errorf("StringCommentInLine(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestCommentMultiLine(t *testing.T) {
	text := strings.Join([]string{
		"lorem ipsum dolor sit amet,",
		"consectetur adipiscing elit.",
		"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	}, "\n")

	expected := strings.Join([]string{
		"/*********************************************************************************",
		" * lorem ipsum dolor sit amet,",
		" * consectetur adipiscing elit.",
		" * sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		" *********************************************************************************/",
	}, "\n")
	got := CommentMultiLine(text)
	if got != expected {
		t.Errorf("StringCommentMultiLine() =\n%q; want\n%q", got, expected)
	}
}

func TestComment(t *testing.T) {
	text1 := "lorem ipsum dolor sit amet"
	expected1 := "/* lorem ipsum dolor sit amet */"
	got1 := Comment(text1)
	if got1 != expected1 {
		t.Errorf("StringComment() = %q; want %q", got1, expected1)
	}

	text2 := strings.Join([]string{
		"lorem ipsum dolor sit amet,",
		"consectetur adipiscing elit.",
		"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	}, "\n")

	expected2 := strings.Join([]string{
		"/*********************************************************************************",
		" * lorem ipsum dolor sit amet,",
		" * consectetur adipiscing elit.",
		" * sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		" *********************************************************************************/",
	}, "\n")
	got2 := Comment(text2)
	if got2 != expected2 {
		t.Errorf("StringComment() =\n%q; want\n%q", got2, expected2)
	}
}
