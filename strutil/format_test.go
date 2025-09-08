package strutil

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "\"\""},
		{"hello", "hello"},
		{"hello world", "\"hello world\""},
	}

	for _, test := range tests {
		result := String(test.input)
		if result != test.expected {
			t.Errorf("StringFormat(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestHexadecimal(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0x0"},
		{0xff, "0xFF"},
		{0x1234abcd, "0x1234ABCD"},
	}

	for _, test := range tests {
		result := Hexadecimal(test.input)
		if result != test.expected {
			t.Errorf("Hexadecimal(%d) = %q; want %q", test.input, result, test.expected)
		}
	}
}
