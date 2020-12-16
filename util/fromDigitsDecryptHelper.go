package util

import "strings"

var dr = strings.NewReplacer(
	"1", "a",
	"2", "e",
	"3", "i",
	"4", "o",
	"5", "u",
)

func ShiftDigitsToLetters(source string) string {

	return dr.Replace(source)

}
