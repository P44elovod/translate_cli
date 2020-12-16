package util

import "strings"

var dr = strings.NewReplacer(
	"1", "a",
	"2", "e",
	"3", "i",
	"4", "o",
	"5", "u",
)

//ShiftDigitsToLetters swap digits to vowel letters in words
func ShiftDigitsToLetters(source string) string {

	return dr.Replace(source)

}
