package keywords

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

func TestGetDataTypeByName(t *testing.T) {
	cases := []struct {
		name     string
		dataType DataType
	}{
		{"UBYTE", UBYTE},
		{"SBYTE", SBYTE},
		{"UWORD", UWORD},
		{"SWORD", SWORD},
		{"ULONG", ULONG},
		{"SLONG", SLONG},
		{"A_UINT64", A_UINT64},
		{"A_INT64", A_INT64},
		{"FLOAT32_IEEE", FLOAT32_IEEE},
		{"FLOAT64_IEEE", FLOAT64_IEEE},
		{"LOREM", InvalidDataType},
	}

	for _, c := range cases {
		gotDt := GetDataTypeByName(c.name)
		if gotDt != c.dataType {
			t.Errorf("GetDataTypeByName(%q) == %s, want %s", c.name, gotDt, c.dataType)
		}

		gotStr := gotDt.String()
		if c.dataType != InvalidDataType && gotStr != c.name {
			t.Errorf("DataType.String() == %q, want %q", gotStr, c.name)
		}
	}
}
