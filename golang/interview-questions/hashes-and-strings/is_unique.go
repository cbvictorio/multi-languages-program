package main

// Implement an algorithm to determine if a string has all unique characters.

func IsUnique(text string) bool {
	dict := make(map[string]bool)

	for _, rune := range text {
		char := string(rune)
		_, found := dict[char]

		if found {
			return false
		}

		dict[char] = true
	}

	return true
}
