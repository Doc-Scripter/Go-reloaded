package inhouse

import (
	"regexp"
)

func Punctuation(s string) string {
	// find  and replace punctuation characters that are not followed by a space
	pattern := regexp.MustCompile(`([.,!?;:])(\S)`)

	s = pattern.ReplaceAllString(s, "$1 $2")
	// find and remove spaces before between characters and puntuation characters
	pattern = regexp.MustCompile(`(\S)(\s)([.,!?;:])`)

	s = pattern.ReplaceAllString(s, "$1$3")

	// find spaces before punctuation characters and replace
	pattern = regexp.MustCompile(`(\s)([.,!?;:])`)

	s = pattern.ReplaceAllString(s, " $2")

	// handle single quotes
	pattern = regexp.MustCompile(`'\s*|\s*'$`)

	s = pattern.ReplaceAllString(s, "'")

	// remove extra spaces between adjacent punctuation characters
	pattern = regexp.MustCompile(`([.,!?;:])\s*([.,!?;:])`)

	s = pattern.ReplaceAllString(s, "$1$2")

	return s
}