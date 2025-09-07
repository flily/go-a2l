package nodes

import (
	"testing"
)

func TestStringFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "\"\""},
		{"hello", "hello"},
		{"hello world", "\"hello world\""},
	}

	for _, test := range tests {
		result := StringFormat(test.input)
		if result != test.expected {
			t.Errorf("StringFormat(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestStringCommentInLine(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"This is a comment.", "/* This is a comment. */"},
	}

	for _, test := range tests {
		result := StringCommentInLine(test.input)
		if result != test.expected {
			t.Errorf("StringCommentInLine(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestTaggedValueStringWith(t *testing.T) {
	tag1 := NewTaggedValue(NoTag, "value1", "value2")
	exp1 := "value1  value2"
	got1 := tag1.String()
	if got1 != exp1 {
		t.Errorf("tag1.String() = %q; want %q", got1, exp1)
	}

	tag2 := NewTaggedValue("TAG", "value1", "value2")
	exp2 := "/* TAG */    value1  value2"
	got2 := tag2.String()
	if got2 != exp2 {
		t.Errorf("tag2.String() = %q; want %q", got2, exp2)
	}

	tag3 := NewTaggedValue("TAG", "value 1", "value 2")
	exp3 := "/* TAG */    \"value 1\"  \"value 2\""
	got3 := tag3.String()
	if got3 != exp3 {
		t.Errorf("tag3.String() = %q; want %q", got3, exp3)
	}
}
