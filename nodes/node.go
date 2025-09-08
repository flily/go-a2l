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
	WriteTo(out io.Writer, indent string, level int, node Node) (int, error)
}

func WriteTo(out io.Writer, indent string, level int, node Node) (int, error) {
	if fn, ok := node.(FormatNode); ok {
		return fn.WriteTo(out, indent, level, node)
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

	n, _ = fmt.Fprintf(out, "%s%s    %s", indentOut, DelimitorEnd, node.Keyword())
	result += n

	return result, nil
}
