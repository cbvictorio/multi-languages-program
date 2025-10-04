package main

import (
	"regexp"
	"strings"
)

type WordData struct {
	TotalCount       int
	UniqueWordsCount int
	Words            []string
}

type WordLengthDictionary map[int]WordData

/*
Some characters must be replaced, others should be deleted.

	firstRegex ->
		Removes quotes, backticks, commas, semicolons, colons,
		exclamation marks, and question marks.

	secondRegex ->
		Replaces dots, newlines (\n), and carriage returns (\r) with spaces.

	thirdRegex ->
		Replaces any remaining special characters
		(anything not alphanumeric, space, or hyphen)
		with spaces.
*/
func ConvertTextToArray(str string) []string {
	str = strings.ReplaceAll(str, "—", "-")

	firstRegex := regexp.MustCompile(`['"\x60,;:!?]`)
	str = firstRegex.ReplaceAllString(str, "")

	secondRegex := regexp.MustCompile(`[.\n\r]`)
	str = secondRegex.ReplaceAllString(str, " ")

	thirdRegex := regexp.MustCompile(`[^a-zA-Z0-9\s-]`)
	str = thirdRegex.ReplaceAllString(str, " ")

	words := strings.Fields(str)

	return words
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

func CreateWordsDictionary(words []string) map[string]int {
	wordsDictionary := make(map[string]int)

	for _, word := range words {
		dictionaryKey := strings.ToLower(word)
		wordCount, exists := wordsDictionary[dictionaryKey]

		if exists {
			wordCount = wordCount + 1
		} else {
			wordCount = 1
		}

		wordsDictionary[dictionaryKey] = wordCount
	}

	return wordsDictionary
}

func CreateWordsLengthMap(words []string) WordLengthDictionary {
	wordLengthDictionary := make(WordLengthDictionary)

	for _, word := range words {
		var newTotalCount int = 1
		var currentUniqueWords []string = []string{word}

		wordLength := len(word)
		value, exists := wordLengthDictionary[wordLength]

		if exists {
			newTotalCount = value.TotalCount + 1
			currentUniqueWords = value.Words

			if !contains(currentUniqueWords, word) {
				currentUniqueWords = append(currentUniqueWords, word)
			}
		}

		wordLengthDictionary[wordLength] = WordData{
			TotalCount:       newTotalCount,
			UniqueWordsCount: len(currentUniqueWords),
			Words:            currentUniqueWords,
		}
	}

	return wordLengthDictionary
}
