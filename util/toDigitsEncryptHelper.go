package util

import (
	"strings"
)

var lr = strings.NewReplacer(
	"a", "1",
	"A", "1",
	"e", "2",
	"E", "2",
	"i", "3",
	"I", "3",
	"o", "4",
	"O", "4",
	"u", "5",
	"U", "5",
)

//ShiftLettersToDigits swap vowel letters to digits in words
func ShiftLettersToDigits(source string) string {

	return lr.Replace(source)

}
