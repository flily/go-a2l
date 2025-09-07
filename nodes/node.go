package nodes

import (
	"fmt"
	"strings"

	"github.com/flily/go-a2l/keywords"
)

type (
	Keyword = keywords.Keyword
)

const (
	DefaultValueDelimiter = "  "
	NoTag                 = ""
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
	Values() []*TaggedValue
	Children() []Node
}
