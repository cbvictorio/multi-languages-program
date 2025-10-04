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

func ConvertStringToArray(str string) []string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9]\s`)
	noSpaceString := regex.ReplaceAllString(str, " ")
	array := strings.Fields(noSpaceString)
	return array
}
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
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
