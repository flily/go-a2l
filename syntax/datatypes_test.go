package syntax

import (
	"testing"
)

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

func TestFloat32ValueString(t *testing.T) {
	v1 := NewFloat32Value(0)
	if v1.ValueString() != "0.000000" {
		t.Errorf("Float32Value(0).String() == %q, want %q", v1.ValueString(), "0.000000")
	}
}
