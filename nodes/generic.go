package nodes

import (
	"fmt"
	"io"

	"github.com/flily/go-a2l/strutil"
	"github.com/flily/go-a2l/syntax"
)

type GenericNodeType uint64

const (
	GenericNodeTypeInline GenericNodeType = 1
	GenericNodeTypeBlock  GenericNodeType = 2
)

type (
	GenericHexadecimalNode struct {
		NodeCommon
		keyword syntax.Keyword
		value   uint64
	}
)

var (
	_ Node = (*GenericHexadecimalNode)(nil)
)

func NewGenericHexadecimalNode(keyword syntax.Keyword, nodeType NodeType, value uint64) *GenericHexadecimalNode {
	n := &GenericHexadecimalNode{
		NodeCommon: NodeCommonInit(keyword, nodeType),
		keyword:    keyword,
		value:      value,
	}

	return n
}

func (n *GenericHexadecimalNode) Values() []*TaggedValue {
	result := []*TaggedValue{
		NewTaggedValue("", strutil.Hexadecimal(n.value)),
	}

	return result
}

func (n *GenericHexadecimalNode) Children() []Node {
	return nil
}

func (n *GenericHexadecimalNode) Comment() string {
	return ""
}

func (n *GenericHexadecimalNode) WriteTo(out io.Writer, indent string, level int) (int, error) {
	return fmt.Fprintf(out, "%s%s        0x%X", indent, n.Keyword(), n.value)
}
