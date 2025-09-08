package strutil

import (
	"fmt"
	"strings"
)

const (
	CommentTitleLine = "********************************************************************************"
)

func CommentInLine(s string) string {
	if len(s) == 0 {
		return ""
	}

	return fmt.Sprintf("/* %s */", s)
}

func CommentMultiLine(s string) string {
	lines := strings.Split(s, "\n")
	result := make([]string, 0, len(lines)+2)

	result = append(result, "/*"+CommentTitleLine)
	for _, line := range lines {
		result = append(result, " * "+line)
	}

	result = append(result, " "+CommentTitleLine+"*/")
	return strings.Join(result, "\n")
}

func Comment(s string) string {
	if strings.Contains(s, "\n") {
		return CommentMultiLine(s)
	}

	return CommentInLine(s)
}
