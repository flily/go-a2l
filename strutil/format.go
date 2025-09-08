package strutil

import (
	"fmt"
	"strings"
)

func String(s string) string {
	if len(s) == 0 {
		return "\"\""
	}

	if strings.Contains(s, " ") {
		return "\"" + s + "\""
	}

	return s
}

func Hexadecimal(n uint64) string {
	return fmt.Sprintf("0x%X", n)
}

func Number[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](n T) string {
	return fmt.Sprintf("%v", n)
}
