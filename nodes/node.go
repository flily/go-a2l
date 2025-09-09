package nodes

import (
	"fmt"
	"io"
	"strings"

	"github.com/flily/go-a2l/strutil"
	"github.com/flily/go-a2l/syntax"
)

type (
	Keyword  = syntax.Keyword
	NodeType int
)

const (
	DefaultValueDelimiter          = "  "
	NoTag                          = ""
	DelimitorBegin                 = "/begin"
	DelimitorEnd                   = "/end"
	NodeTypeTopLevel      NodeType = 1
	NodeTypePrimary       NodeType = 2
	NodeTypeSecondary     NodeType = 3
	NodeTypeAttribute     NodeType = 4
)

type TaggedValue struct {
	Tag    string
	Values []string
}

func NewTaggedValue(tag string, values ...string) *TaggedValue {
	v := &TaggedValue{
		Tag:    tag,
		Values: values,
	}

	return v
}

func (v *TaggedValue) GetTag() string {
	tag := ""
	if len(v.Tag) > 0 {
		tag = strutil.CommentInLine(v.Tag)
	}

	return tag
}

func (v *TaggedValue) GetValues(delim string) string {
	result := make([]string, len(v.Values))
	for i, val := range v.Values {
		result[i] = strutil.String(val)
	}

	return strings.Join(result, delim)
}

func (v *TaggedValue) Content() (string, string) {
	return v.GetTag(), v.GetValues(DefaultValueDelimiter)
}

func (v *TaggedValue) String() string {
	tag, values := v.Content()
	if len(tag) <= 0 {
		return values
	}

	return fmt.Sprintf("%s    %s", tag, values)
}

type Node interface {
	Keyword() Keyword
	NodeType() NodeType
	Values() []*TaggedValue
	Children() []Node
	Comment() string
}

type FormatNode interface {
	WriteTo(out io.Writer, indent string, level int) (int, error)
}

func WriteTo(out io.Writer, indent string, level int, node Node) (int, error) {
	if fn, ok := node.(FormatNode); ok {
		return fn.WriteTo(out, indent, level)
	}

	n := 0
	result := 0

	indentOut := strings.Repeat(indent, level)
	//                       /begin  NAME
	//                       /end    NAME
	n, _ = fmt.Fprintf(out, "%s%s  %s\n", indentOut, DelimitorBegin, node.Keyword())
	result += n

	indentIn := strings.Repeat(indent, level+1)

	fields := node.Values()
	for i := 0; i < len(fields); i++ {
		comment, value := fields[i].Content()
		n, _ = fmt.Fprintf(out, "%s%s        %s\n", indentIn, comment, value)
		result += n
	}

	children := node.Children()
	for i := 0; i < len(children); i++ {
		n, err := WriteTo(out, indent, level+1, children[i])
		if err != nil {
			return 0, err
		}

		_, _ = fmt.Fprint(out, "\n")
		result += n + 1
	}

	n, _ = fmt.Fprintf(out, "%s%s    %s", indentOut, DelimitorEnd, node.Keyword())
	result += n

	return result, nil
}

func WriteString(indent string, level int, node Node) (string, error) {
	var sb strings.Builder
	_, err := WriteTo(&sb, indent, level, node)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

type NodeCommon struct {
	keyword  syntax.Keyword
	nodeType NodeType
}

func (n *NodeCommon) Keyword() Keyword {
	return n.keyword
}

func (n *NodeCommon) NodeType() NodeType {
	return n.nodeType
}

func NodeCommonInit(keyword syntax.Keyword, nodeType NodeType) NodeCommon {
	n := NodeCommon{
		keyword:  keyword,
		nodeType: nodeType,
	}

	return n
}

func PrimaryNodeInit(keyword syntax.Keyword) NodeCommon {
	return NodeCommonInit(keyword, NodeTypePrimary)
}

func SecondaryNodeInit(keyword syntax.Keyword) NodeCommon {
	return NodeCommonInit(keyword, NodeTypeSecondary)
}
