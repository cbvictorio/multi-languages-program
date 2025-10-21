package main

import (
	"sort"
)

func SortString(text string) string {
	textToRunes := []rune(text)

	sort.Slice(textToRunes, func(i, j int) bool {
		return textToRunes[i] < textToRunes[j]
	})

	return string(textToRunes)
}

func CheckPermutation(firstText string, secondText string) bool {

	sortedFirstText := SortString(firstText)
	loopLimit := len(secondText) - len(firstText)

	for i := 0; i <= loopLimit; i++ {
		indexStep := i + len(firstText)
		subString := secondText[i:indexStep]
		sortedSubString := SortString(subString)

		if sortedFirstText == sortedSubString {
			return true
		}
	}

	return false
}
