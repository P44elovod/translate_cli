package util

import "strings"

var vowels = "aeiouy"

func TranslateWordToPiglatin(sourceWord string) string {
	firstVowelIndex := getFirstVowelIndex(sourceWord)
	if firstVowelIndex == 0 {

		return sourceWord + "yay"
	}

	return sourceWord[firstVowelIndex:] + sourceWord[:firstVowelIndex] + "ay"
}

func getFirstVowelIndex(str string) int {
	return strings.IndexAny(str, vowels)
}
