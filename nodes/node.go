package nodes

import (
	"fmt"
	"io"
	"strings"

	"github.com/flily/go-a2l/keywords"
)

type (
	Keyword  = keywords.Keyword
	NodeType int
)

const (
	DefaultValueDelimiter          = "  "
	NoTag                          = ""
	DelimitorBegin                 = "/begin"
	DelimitorEnd                   = "/end"
	CommentTitleLine               = "********************************************************************************"
	NodeTypeTopLevel      NodeType = 1
	NodeTypePrimary       NodeType = 2
	NodeTypeSecondary     NodeType = 3
	NodeTypeAttribute     NodeType = 4
)

func StringFormat(s string) string {
	if len(s) == 0 {
		return "\"\""
	}

	if strings.Contains(s, " ") {
		return "\"" + s + "\""
	}

	return s
}

func StringCommentInLine(s string) string {
	if len(s) == 0 {
		return ""
	}

	return fmt.Sprintf("/* %s */", s)
}

func StringCommentMultiLine(s string) string {
	lines := strings.Split(s, "\n")
	result := make([]string, 0, len(lines)+2)

	result = append(result, "/*"+CommentTitleLine)
	for _, line := range lines {
		result = append(result, " * "+line)
	}

	result = append(result, " "+CommentTitleLine+"*/")
	return strings.Join(result, "\n")
}

func StringComment(s string) string {
	if strings.Contains(s, "\n") {
		return StringCommentMultiLine(s)
	}

	return StringCommentInLine(s)
}

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
		tag = StringCommentInLine(v.Tag)
	}

	return tag
}

func (v *TaggedValue) GetValues(delim string) string {
	result := make([]string, len(v.Values))
	for i, val := range v.Values {
		result[i] = StringFormat(val)
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
	Name() Keyword
	Type() NodeType
	Values() []*TaggedValue
	Children() []Node
	Comment() string
	WriteString(indent string) string
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
	n, _ = fmt.Fprintf(out, "%s%s  %s\n", indentOut, DelimitorBegin, node.Name())
	result += n

	indentIn := strings.Repeat(indent, level+1)

	fields := node.Values()
	for i := 0; i < len(fields); i++ {
		comment, value := fields[i].Content()
		n, _ = fmt.Fprintf(out, "%s%s     %s\n", indentIn, StringCommentInLine(comment), StringFormat(value))
		result += n
	}

	n, _ = fmt.Fprintf(out, "%s%s     %s\n", indentOut, DelimitorEnd, node.Name())
	result += n

	return result, nil
}
