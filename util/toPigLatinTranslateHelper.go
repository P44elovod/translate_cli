package util

import (
	"regexp"
	"strings"
)

var vowels = "aeiouy"

//TranslateTextToPiglatin is used for translation fron English to Piglatin
func TranslateTextToPiglatin(sourceText string) string {

	words := strings.Split(regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(sourceText, ""), " ")

	for i := 0; i < len(words); i++ {
		sourceText = regexp.MustCompile(`\b`+words[i]+`\b`).ReplaceAllString(sourceText, translateWordToPiglatin(words[i]))
	}

	return sourceText
}

func translateWordToPiglatin(sourceWord string) string {
	firstVowelIndex := getFirstVowelIndex(sourceWord)

	if firstVowelIndex == -1 {
		return sourceWord + "(this word couldn't be translated)"
	}

	if firstVowelIndex != 0 && len(sourceWord) > 1 {
		return sourceWord[firstVowelIndex:] + sourceWord[:firstVowelIndex] + "ay"
	}

	return sourceWord + "yay"
}

func getFirstVowelIndex(str string) int {
	return strings.IndexAny(str, vowels)
}
