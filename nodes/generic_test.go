package nodes

import (
	"testing"

	"github.com/flily/go-a2l/syntax"
)

func TestGenericHexadecialNodeWriteString(t *testing.T) {
	node := NewGenericHexadecimalNode(syntax.ECU_ADDRESS, NodeTypePrimary, 0x1a2b3c4d)

	expected := "  ECU_ADDRESS        0x1A2B3C4D"
	got, err := WriteString("  ", 0, node)
	if err != nil {
		t.Fatalf("WriteTo() error: %v", err)
	}

	if got != expected {
		t.Errorf("GenericHexadecimalNode.WriteTo() =\n%s; want\n%s", got, expected)
	}
}
