package utils

import (
	"strings"
	"unicode"
)

func Parse(toParse string) string {
	var message = strings.Split(toParse, ": ")[1]
	return capitalize(message)
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
