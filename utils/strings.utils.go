package utils

import (
	"strings"
	"unicode"
)

func ToCapital(str string) string {
	b := []byte(strings.ToLower(str))
	b[0] = byte(unicode.ToUpper(rune(b[0])))
	capital := string(b)
	return capital
}
