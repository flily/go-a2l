package syntax

import (
	"testing"
)

func TestGetKeywordByName(t *testing.T) {
	cases := []struct {
		name    string
		keyword Keyword
	}{
		{"PROJECT", PROJECT},
		{"MODULE", MODULE},
		{"A2ML", A2ML},
		{"MEASUREMENT", MEASUREMENT},
		{"FORMAT", FORMAT},
		{"FORMULA", FORMULA},
		{"LOREM", InvalidKeyword},
	}

	for _, c := range cases {
		gotKw := GetKeywordByName(c.name)
		if gotKw != c.keyword {
			t.Errorf("GetKeywordByName(%q) == %s, want %s", c.name, gotKw, c.keyword)
		}

		gotStr := gotKw.String()
		if c.keyword != InvalidKeyword && gotStr != c.name {
			t.Errorf("Keyword.String() == %q, want %q", gotStr, c.name)
		}
	}
}
