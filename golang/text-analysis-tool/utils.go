package main

import (
	"strings"
)

// type WordLengthDictionary = {
// 	totalCount: number;
// 	uniqueWordsCount: number;
// 	words: string[];
//   };

type WordData struct {
	TotalCount       int
	UniqueWordsCount int
	Words            []string
}

type WordLengthDictionary map[int]WordData

func ConvertStringToArray(str string) []string {
	array := strings.Fields(str)
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
		var currentUniqueWords []string

		wordLength := len(word)
		value, exists := wordLengthDictionary[wordLength]

		if exists {
			newTotalCount = value.TotalCount + 1
			currentUniqueWords = value.Words

			if !contains(currentUniqueWords, word) {
				currentUniqueWords = append(currentUniqueWords, word)
			}
		} else {
			currentUniqueWords = []string{word}
		}

		wordLengthDictionary[wordLength] = WordData{
			TotalCount:       newTotalCount,
			UniqueWordsCount: len(currentUniqueWords),
			Words:            currentUniqueWords,
		}
	}

	return wordLengthDictionary
}
